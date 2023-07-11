# 🚀 go-semantic-versioning
Version packages with the Semantic Versioning standard, and Go.

# 🍺 Brew
brew tap wpartin/go-semantic-versioning && brew install go-semantic-versioning

# 📃 Philosophy
The idea behind this project is to have a clean, trunk-based repository for your application, while being able to version specific points easily. You can plug this tool into your CI/CD pipeline, and based upon the branch prefix merging into main/master, create an incremented version in your Git history.

# 🔍 Example
Create feature/* branch off of master -> make code changes -> create a Pull Request for review -> merge kicks off a CI/CD tool that you use -> that tool runs this one as a step -> the tool creates your tag and applies it to that point of the repository

# 🎁 Extras
If you have goreleaser installed, you can use the -release flag to create your release on GitHub.
