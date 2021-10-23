# Speecurve CLI

A tool to automate usage of speedcurve.com. For now it is intended to be used from your local machine to trigger deploy tests and get latest deploys instead of using the GUI.

# How to use
- Rename .env.sample file to .env
- Update .env file to include your speedcurve API_KEY

## To get deploys
- `cd get-deploy`
- `go run .`

## To add deploy
- `cd add-deploy`
- `go run .`
- Follow onscreen instructions to select the site, short note and description