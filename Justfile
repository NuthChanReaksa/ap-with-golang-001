
build:
    # Add your build commands here

clean:
    # Add your clean commands here

run:
    # Add your run commands here
    go run cmd/app/main.go


branch := `git branch --show-current`

autopush:
    @echo "Checking git status..."
    git status

    @echo "Checking for changes to commit..."
    @sh -c 'if git diff-index --quiet HEAD --; then \
        echo "No changes to commit."; \
    else \
        echo "Adding all changes..."; \
        git add .; \
        echo "Enter commit message: "; \
        read msg; \
        echo "Committing changes with message: $$msg"; \
        git commit -m "$$msg"; \
        echo "Pushing to branch {{branch}}..."; \
        git push origin {{branch}}; \
        echo "Push complete!"; \
    fi'

