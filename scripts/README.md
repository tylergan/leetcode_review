# Scripts

Utility scripts for managing the LeetcodeReview repository.

## update_readme.py

Automatically generates or updates the main README.md by scanning all problem directories.

### Features:
- **Auto-detects categories**: Scans all folders and treats them as problem categories
- **Configurable**: Uses `.leetcode-config.yml` to ignore specific folders
- **GitHub Pages support**: Automatically generates correct URLs for hosted cheatsheets
- **Smart detection**: Only includes folders that contain Go problem files

### Usage:

```bash
# Run manually
python3 scripts/update_readme.py

# Or make it executable and run directly
./scripts/update_readme.py
```

### Configuration:

Edit `.leetcode-config.yml` in the project root to customize:

```yaml
# Folders to ignore when scanning for problem categories
ignore_folders:
  - .git
  - scripts
  - docs

# GitHub Pages settings
github_pages:
  enabled: true
  username: your-github-username
  repository: your-repo-name
```

### Automatic Updates:

The pre-commit hook automatically runs this script before each commit, ensuring the README is always up-to-date.
