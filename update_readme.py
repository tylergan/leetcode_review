#!/usr/bin/env python3
"""
Generate or update README.md for the LeetcodeReview project.
Scans directories for problem files and creates organized documentation.
"""

import os
import re
from pathlib import Path
from typing import List, Dict, Tuple

# ============================================================================
# GitHub Pages Configuration
# ============================================================================
# If you have GitHub Pages enabled, set USE_GITHUB_PAGES to True and update
# your GitHub username and repository name below. The cheatsheet link will
# use the GitHub Pages URL (https://username.github.io/repo/...).
#
# If you prefer local file paths, set USE_GITHUB_PAGES to False.
# ============================================================================
GITHUB_USERNAME = "tylergan"  # Change this to your GitHub username
GITHUB_REPO = "leetcode_review"  # Change this to your repo name
USE_GITHUB_PAGES = True  # Set to False to use local paths


def extract_problem_info(file_path: Path) -> Dict[str, str]:
    """Extract problem name and description from a Go file."""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Extract from comment block
    comment_match = re.search(r'/\*\s*(.*?)\s*\*/', content, re.DOTALL)
    if not comment_match:
        return {}

    comment_text = comment_match.group(1)

    # Extract goal/description (first non-empty line usually)
    lines = [line.strip() for line in comment_text.split('\n') if line.strip()]
    description = lines[0] if lines else ""

    # Extract example if present
    example_match = re.search(r'Example \d+:(.*?)(?=Example \d+:|Output:|Constraints:|$)',
                             comment_text, re.DOTALL)
    example = example_match.group(1).strip() if example_match else ""

    return {
        'description': description,
        'example': example
    }


def get_problem_title(filename: str) -> str:
    """Convert filename to readable title."""
    # Remove .go extension and convert underscores to spaces
    title = filename.replace('.go', '').replace('_', ' ').title()
    return title


def scan_directory(dir_path: Path) -> List[Tuple[str, Path, Dict]]:
    """Scan directory for problem files and extract info."""
    problems = []

    for file_path in sorted(dir_path.glob('*.go')):
        # Skip test files
        if file_path.name.endswith('_test.go'):
            continue

        info = extract_problem_info(file_path)
        title = get_problem_title(file_path.name)
        problems.append((title, file_path, info))

    return problems


def generate_readme(base_path: Path) -> str:
    """Generate complete README content."""
    readme_lines = [
        "# LeetCode Review",
        "",
        "A collection of LeetCode problems solved in Go with comprehensive test coverage.",
        "",
        "## 📚 Problem Categories",
        ""
    ]

    # Scan each category directory
    categories = ['arrays']  # Add more categories as needed

    for category in categories:
        category_path = base_path / category
        if not category_path.exists():
            continue

        # Add category header
        readme_lines.append(f"### {category.title()}")
        readme_lines.append("")

        # Check if cheatsheet exists
        cheatsheet_path = category_path / 'cheatsheet.html'
        if cheatsheet_path.exists():
            if USE_GITHUB_PAGES:
                cheatsheet_url = f"https://{GITHUB_USERNAME}.github.io/{GITHUB_REPO}/{category}/cheatsheet.html"
            else:
                cheatsheet_url = f"{category}/cheatsheet.html"
            readme_lines.append(f"📄 **[View Cheat Sheet]({cheatsheet_url})** - Quick reference guide for {category} problems")
            readme_lines.append("")

        # Scan for problems
        problems = scan_directory(category_path)

        if problems:
            readme_lines.append("| Problem | Description | Solution | Tests |")
            readme_lines.append("|---------|-------------|----------|-------|")

            for title, file_path, info in problems:
                rel_path = file_path.relative_to(base_path)
                test_file = file_path.parent / f"{file_path.stem}_test.go"

                description = info.get('description', '')[:80]
                if len(info.get('description', '')) > 80:
                    description += "..."

                solution_link = f"[{file_path.name}]({rel_path})"

                if test_file.exists():
                    test_rel_path = test_file.relative_to(base_path)
                    test_link = f"[✓]({test_rel_path})"
                else:
                    test_link = "❌"

                readme_lines.append(f"| {title} | {description} | {solution_link} | {test_link} |")

            readme_lines.append("")

    # Add usage section
    readme_lines.extend([
        "## 🚀 Running Tests",
        "",
        "Run all tests:",
        "```bash",
        "go test ./...",
        "```",
        "",
        "Run tests for a specific category:",
        "```bash",
        "go test ./arrays",
        "```",
        "",
        "Run tests for a specific problem:",
        "```bash",
        "go test ./arrays -run TestTwoSum",
        "```",
        "",
        "## 🎯 Pre-commit Hook",
        "",
        "This repository includes a pre-commit hook that:",
        "- Formats Go code with `go fmt`",
        "- Runs all tests before allowing commits",
        "",
        "The hook is automatically set up in `.git/hooks/pre-commit`.",
        "",
        "## 📖 Cheat Sheets",
        "",
        "Each category includes a comprehensive cheat sheet with:",
        "- Common patterns and techniques",
        "- Time and space complexity analysis",
        "- Key insights and gotchas",
        "- Code snippets",
        ""
    ])

    if USE_GITHUB_PAGES:
        readme_lines.append(f"Cheat sheets are hosted on GitHub Pages. Click the links above to view them online!")
    else:
        readme_lines.append("View the cheat sheets by opening the HTML files in your browser.")

    readme_lines.extend([
        "",
        "## 🔄 Updating This README",
        "",
        "This README is automatically generated. To update it, run:",
        "```bash",
        "python3 update_readme.py",
        "```",
        ""
    ])

    return "\n".join(readme_lines)


def main():
    """Main entry point."""
    # Get the base directory (where the script is located)
    script_dir = Path(__file__).parent

    # Generate README content
    readme_content = generate_readme(script_dir)

    # Write to README.md
    readme_path = script_dir / 'README.md'
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(readme_content)

    print(f"✓ README.md updated successfully!")
    print(f"  Location: {readme_path}")


if __name__ == '__main__':
    main()
