package graph

import (
	"testing"
)

func TestValidateGraph(t *testing.T) {
	mockConfig := Manifest{
		Scopes:    []string{"definition", "description", "explanation", "lesson"},
		Clarities: []string{"vague", "introductory", "detailed", "strict"},
	}
	tests := []struct {
		name    string
		guides  map[string]Guide
		wantErr bool
	}{
		{
			name: "valid single guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   "definition",
						Clarity: "strict",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid scope",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   "invalid_scope",
						Clarity: "strict",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid clarity",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   "definition",
						Clarity: "invalid_clarity",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid prerequisite",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   "lesson",
						Clarity: "detailed",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unknown prerequisite",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"unknown_guide"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "mismatched prerequisite scope",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   "definition",
						Clarity: "detailed",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid sub-guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     "lesson",
						Clarity:   "detailed",
						SubGuides: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   "explanation",
						Clarity: "detailed",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unknown sub-guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     "lesson",
						Clarity:   "detailed",
						SubGuides: []string{"unknown_guide"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid sub-guide scope (jumping level without relaxed_subguides)",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     "lesson", // 4
						Clarity:   "detailed",
						SubGuides: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   "description", // 2 (differs by 2)
						Clarity: "detailed",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid sub-guide scope (larger scope)",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     "explanation",
						Clarity:   "detailed",
						SubGuides: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   "lesson", // larger scope than Explanation
						Clarity: "detailed",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "cyclic prerequisites",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide1"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "transitive cycle in prerequisites",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide3"},
					},
				},
				"guide3": {
					ID: "guide3",
					Metadata: GuideMetadata{
						Scope:         "lesson",
						Clarity:       "detailed",
						Prerequisites: []string{"guide1"},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGraph(tt.guides, mockConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGraph() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckAcyclic(t *testing.T) {
	tests := []struct {
		name    string
		guides  map[string]Guide
		wantErr bool
	}{
		{
			name:    "empty graph",
			guides:  map[string]Guide{},
			wantErr: false,
		},
		{
			name: "single node without dependencies",
			guides: map[string]Guide{
				"a": {ID: "a"},
			},
			wantErr: false,
		},
		{
			name: "linear acyclic chain",
			guides: map[string]Guide{
				"a": {ID: "a", Metadata: GuideMetadata{Prerequisites: []string{"b"}}},
				"b": {ID: "b", Metadata: GuideMetadata{Prerequisites: []string{"c"}}},
				"c": {ID: "c"},
			},
			wantErr: false,
		},
		{
			name: "diamond acyclic graph",
			guides: map[string]Guide{
				"a": {ID: "a", Metadata: GuideMetadata{Prerequisites: []string{"b", "c"}}},
				"b": {ID: "b", Metadata: GuideMetadata{Prerequisites: []string{"d"}}},
				"c": {ID: "c", Metadata: GuideMetadata{Prerequisites: []string{"d"}}},
				"d": {ID: "d"},
			},
			wantErr: false,
		},
		{
			name: "self-referential cycle",
			guides: map[string]Guide{
				"a": {ID: "a", Metadata: GuideMetadata{Prerequisites: []string{"a"}}},
			},
			wantErr: true,
		},
		{
			name: "sub_guide cycle",
			guides: map[string]Guide{
				"a": {ID: "a", Metadata: GuideMetadata{SubGuides: []string{"b"}}},
				"b": {ID: "b", Metadata: GuideMetadata{SubGuides: []string{"a"}}},
			},
			wantErr: true,
		},
		{
			name: "mixed prerequisites and sub_guides cycle",
			guides: map[string]Guide{
				"a": {ID: "a", Metadata: GuideMetadata{Prerequisites: []string{"b"}}},
				"b": {ID: "b", Metadata: GuideMetadata{SubGuides: []string{"c"}}},
				"c": {ID: "c", Metadata: GuideMetadata{Prerequisites: []string{"a"}}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckAcyclic(tt.guides)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAcyclic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateGraphRelaxedSubguides(t *testing.T) {
	mockConfig := Manifest{
		Scopes:           []string{"definition", "description", "explanation", "lesson"},
		Clarities:        []string{"vague", "introductory", "detailed", "strict"},
		RelaxedSubguides: true,
	}

	guides := map[string]Guide{
		"guide1": {
			ID: "guide1",
			Metadata: GuideMetadata{
				Scope:     "lesson", // 4
				Clarity:   "detailed",
				SubGuides: []string{"guide2"},
			},
		},
		"guide2": {
			ID: "guide2",
			Metadata: GuideMetadata{
				Scope:   "description", // 2 (differs by 2)
				Clarity: "detailed",
			},
		},
	}

	err := ValidateGraph(guides, mockConfig)
	if err != nil {
		t.Errorf("ValidateGraph() with relaxed_subguides=true returned error = %v, want no error", err)
	}
}
