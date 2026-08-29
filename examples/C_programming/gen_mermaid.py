import os
import yaml

def generate_mermaid():
    base_dir = "."
    files = [f for f in os.listdir(base_dir) if f.endswith(".md")]
    
    nodes = {}
    edges = []
    
    # First, read all guides
    for f in files:
        path = os.path.join(base_dir, f)
        with open(path, "r") as file:
            content = file.read()
            parts = content.split("---")
            if len(parts) >= 3:
                fm = parts[1]
                data = yaml.safe_load(fm)
                if data:
                    name = f.replace(".md", "")
                    nodes[name] = data
                    
    # Generate edges
    for name, data in nodes.items():
        prereqs = data.get("prerequisites", [])
        if prereqs:
            for p in prereqs:
                edges.append(f'    "{p}" --> "{name}"')
                
        sub_guides = data.get("sub_guides", [])
        if sub_guides:
            for sg in sub_guides:
                if isinstance(sg, dict):
                    sg_name = sg.get("guide")
                else:
                    sg_name = sg
                if sg_name:
                    edges.append(f'    "{name}" -.->|sub| "{sg_name}"')

    # Output mermaid file
    with open("graph_diagram.mmd", "w") as out:
        out.write("graph TD\n")
        out.write("    %% Prerequisites (solid arrows)\n")
        out.write("    %% Sub-guides (dotted arrows)\n\n")
        for edge in edges:
            out.write(edge + "\n")
            
    print("Mermaid diagram generated successfully.")

if __name__ == '__main__':
    generate_mermaid()
