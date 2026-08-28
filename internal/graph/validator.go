package graph

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// CheckAcyclic checks if the knowledge graph is acyclic using DFS
func CheckAcyclic(guides map[string]Guide) error {
	const (
		unvisited = 0
		visiting  = 1
		visited   = 2
	)

	state := make(map[string]int, len(guides))

	var dfs func(id string) error
	dfs = func(id string) error {
		state[id] = visiting

		if guide, exists := guides[id]; exists {
			neighbors := append([]string(nil), guide.Metadata.Prerequisites...)
			for _, sub := range guide.Metadata.SubGuides {
				neighbors = append(neighbors, sub.Guide)
			}

			for _, neighborID := range neighbors {
				if _, exists := guides[neighborID]; !exists {
					continue
				}
				if state[neighborID] == visiting {
					return fmt.Errorf("cycle detected in graph involving guide '%s'", neighborID)
				}
				if state[neighborID] == unvisited {
					if err := dfs(neighborID); err != nil {
						return err
					}
				}
			}
		}

		state[id] = visited
		return nil
	}

	for id := range guides {
		if state[id] == unvisited {
			if err := dfs(id); err != nil {
				return err
			}
		}
	}

	return nil
}

// parseSegments parses a segment string like "10-20, 25-30" into a slice of [2]int
func parseSegments(segmentStr string) ([][2]int, error) {
	if strings.TrimSpace(segmentStr) == "" {
		return nil, nil
	}
	var ranges [][2]int
	parts := strings.Split(segmentStr, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			return nil, fmt.Errorf("invalid range format: %s", p)
		}
		start, err := strconv.Atoi(strings.TrimSpace(bounds[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start line: %w", err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(bounds[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end line: %w", err)
		}
		if start > end {
			return nil, fmt.Errorf("start line %d is greater than end line %d", start, end)
		}
		ranges = append(ranges, [2]int{start, end})
	}
	return ranges, nil
}

// getScopeValue helper returns the hierarchical index of the given scope (1-indexed)
func getScopeValue(scope string, config Manifest) int {
	for i, s := range config.Scopes {
		if s == scope {
			return i + 1
		}
	}
	return 0
}

func isValidClarity(clarity string, config Manifest) bool {
	return slices.Contains(config.Clarities, clarity)
}

// ValidateGraph checks structural constraints of the knowledge graph
func ValidateGraph(guides map[string]Guide, config Manifest) error {
	for _, guide := range guides {
		// Validate Scope
		if getScopeValue(guide.Metadata.Scope, config) == 0 {
			return fmt.Errorf("guide '%s' has invalid scope: '%s'", guide.ID, guide.Metadata.Scope)
		}

		// Validate Clarity
		if !isValidClarity(guide.Metadata.Clarity, config) {
			return fmt.Errorf("guide '%s' has invalid clarity: '%s'", guide.ID, guide.Metadata.Clarity)
		}

		// Validate Prerequisites
		for _, prereqID := range guide.Metadata.Prerequisites {
			prereq, exists := guides[prereqID]
			if !exists {
				return fmt.Errorf("guide '%s' references unknown prerequisite: '%s'", guide.ID, prereqID)
			}
			if prereq.Metadata.Scope != guide.Metadata.Scope {
				return fmt.Errorf("guide '%s' (scope: '%s') has prerequisite '%s' with mismatched scope: '%s'. Horizontal edges must have exactly identical scope",
					guide.ID, guide.Metadata.Scope, prereqID, prereq.Metadata.Scope)
			}
		}

		var allRanges [][2]int

		// Validate SubGuides
		for _, subRelation := range guide.Metadata.SubGuides {
			subID := subRelation.Guide
			sub, exists := guides[subID]
			if !exists {
				return fmt.Errorf("guide '%s' references unknown sub_guide: '%s'", guide.ID, subID)
			}

			if subRelation.Clarity != "" && !isValidClarity(subRelation.Clarity, config) {
				return fmt.Errorf("guide '%s' has invalid subguide clarity: '%s' for subguide '%s'", guide.ID, subRelation.Clarity, subID)
			}

			if subRelation.Segment != "" {
				ranges, err := parseSegments(subRelation.Segment)
				if err != nil {
					return fmt.Errorf("guide '%s' has invalid segment for subguide '%s': %w", guide.ID, subID, err)
				}
				allRanges = append(allRanges, ranges...)
			}

			guideScopeVal := getScopeValue(guide.Metadata.Scope, config)
			subScopeVal := getScopeValue(sub.Metadata.Scope, config)

			if subScopeVal >= guideScopeVal {
				return fmt.Errorf("guide '%s' (scope: '%s') has sub_guide '%s' with invalid scope: '%s'. Sub-guides must have a strictly smaller scope",
					guide.ID, guide.Metadata.Scope, subID, sub.Metadata.Scope)
			}

			if !config.RelaxedSubguides && guideScopeVal-subScopeVal != 1 {
				return fmt.Errorf("guide '%s' (scope: '%s') has sub_guide '%s' with scope: '%s', but relaxed_subguides is false. Sub-guides must be exactly one scope level below their parent",
					guide.ID, guide.Metadata.Scope, subID, sub.Metadata.Scope)
			}
		}

		// Sort and check overlaps
		slices.SortFunc(allRanges, func(a, b [2]int) int {
			if a[0] == b[0] {
				return a[1] - b[1]
			}
			return a[0] - b[0]
		})

		for i := 1; i < len(allRanges); i++ {
			if allRanges[i][0] <= allRanges[i-1][1] {
				return fmt.Errorf("guide '%s' has overlapping subguide segments: %v and %v", guide.ID, allRanges[i-1], allRanges[i])
			}
		}

		if config.StrictCoverage && len(guide.Metadata.SubGuides) > 0 {
			if len(allRanges) == 0 {
				return fmt.Errorf("guide '%s' has strict_coverage enabled but no segments defined", guide.ID)
			}
			if allRanges[0][0] > 1 {
				return fmt.Errorf("guide '%s' is missing coverage for lines 1 to %d", guide.ID, allRanges[0][0]-1)
			}
			for i := 1; i < len(allRanges); i++ {
				if allRanges[i][0] > allRanges[i-1][1]+1 {
					return fmt.Errorf("guide '%s' is missing coverage for lines %d to %d", guide.ID, allRanges[i-1][1]+1, allRanges[i][0]-1)
				}
			}
			if allRanges[len(allRanges)-1][1] < guide.LineCount {
				return fmt.Errorf("guide '%s' is missing coverage for lines %d to %d", guide.ID, allRanges[len(allRanges)-1][1]+1, guide.LineCount)
			}
		}
	}
	return CheckAcyclic(guides)
}
