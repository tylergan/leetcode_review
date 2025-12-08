#!/usr/bin/env python3
"""
Generate or update README.md for the LeetcodeReview project.
Scans directories for problem files and creates organized documentation.
"""

import os
import re
import yaml
from pathlib import Path
from typing import List, Dict, Tuple, Set


def load_config(base_path: Path) -> Dict:
    """Load configuration from .leetcode-config.yml"""
    config_path = base_path / '.leetcode-config.yml'

    # Default configuration
    default_config = {
        'ignore_folders': ['.git', '.github', 'scripts', 'docs', 'node_modules', 'venv', '__pycache__'],
        'github_pages': {
            'enabled': True,
            'username': 'tylergan',
            'repository': 'leetcode_review'
        }
    }

    if not config_path.exists():
        print(f"⚠️  No config file found at {config_path}, using defaults")
        return default_config

    try:
        with open(config_path, 'r', encoding='utf-8') as f:
            config = yaml.safe_load(f)
            return config if config else default_config
    except Exception as e:
        print(f"⚠️  Error loading config: {e}, using defaults")
        return default_config


def get_categories(base_path: Path, ignore_folders: Set[str]) -> List[str]:
    """Auto-detect problem categories by scanning directories."""
    categories = []

    for item in sorted(base_path.iterdir()):
        # Only consider directories
        if not item.is_dir():
            continue

        # Skip hidden folders and ignored folders
        if item.name.startswith('.') or item.name in ignore_folders:
            continue

        # Check if directory contains any .go files (excluding tests)
        has_problems = any(
            f.suffix == '.go' and not f.name.endswith('_test.go')
            for f in item.glob('*.go')
        )

        if has_problems:
            categories.append(item.name)

    return categories


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


def generate_readme(base_path: Path, config: Dict) -> str:
    """Generate complete README content."""
    readme_lines = [
        "# LeetCode Review",
        "",
        "A collection of LeetCode problems solved in Go with comprehensive test coverage.",
        "",
        "## 📚 Problem Categories",
        ""
    ]

    # Get configuration
    ignore_folders = set(config.get('ignore_folders', []))
    github_pages = config.get('github_pages', {})
    use_github_pages = github_pages.get('enabled', False)
    github_username = github_pages.get('username', '')
    github_repo = github_pages.get('repository', '')

    # Auto-detect categories
    categories = get_categories(base_path, ignore_folders)

    if not categories:
        readme_lines.append("_No problem categories found yet. Add your first problem to get started!_")
        readme_lines.append("")

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
            if use_github_pages and github_username and github_repo:
                cheatsheet_url = f"https://{github_username}.github.io/{github_repo}/{category}/cheatsheet.html"
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
        "- Updates this README automatically",
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

    if use_github_pages:
        readme_lines.append(f"Cheat sheets are hosted on GitHub Pages. Click the links above to view them online!")
    else:
        readme_lines.append("View the cheat sheets by opening the HTML files in your browser.")

    readme_lines.extend([
        "",
        "## ⚙️ Configuration",
        "",
        "Customize the README generator by editing `.leetcode-config.yml`:",
        "- Add folders to ignore when scanning for categories",
        "- Configure GitHub Pages settings",
        "",
        "## 🔄 Updating This README",
        "",
        "This README is automatically generated. To update it manually, run:",
        "```bash",
        "python3 scripts/update_readme.py",
        "```",
        ""
    ])

    return "\n".join(readme_lines)


def main():
    """Main entry point."""
    # Get the base directory (parent of scripts folder where the script is located)
    script_dir = Path(__file__).parent
    base_path = script_dir.parent

    # Load configuration
    config = load_config(base_path)

    # Generate README content
    readme_content = generate_readme(base_path, config)

    # Write to README.md
    readme_path = base_path / 'README.md'
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(readme_content)

    print(f"✓ README.md updated successfully!")
    print(f"  Location: {readme_path}")

    # Show detected categories
    ignore_folders = set(config.get('ignore_folders', []))
    categories = get_categories(base_path, ignore_folders)
    if categories:
        print(f"  Categories: {', '.join(categories)}")


if __name__ == '__main__':
    main()
