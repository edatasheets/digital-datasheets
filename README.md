# Digital Datasheet Properties

This repository contains information about the format and properties
of digital datasheets for a variety of component types.

Document types include:
* Specification for component properties
* Specification for pin properties
* Sample .json format for different types of digital datasheets

# Contributing

If you are a working group member, to commit changes, submit a PR and request a 
review from at least one working group member.

If you are not a working group memeber, submit an issue to discuss your proposed
change before submitting a PR.

# Releasing

A new version of digital datasheets is officially released by merging a PR with 
changes for that version into the release repo: <https://github.com/edatasheets/edatasheets.github>

The release.sh script automates copying the appropriate files, generating a new
README.md, and updating links. The script requires a version of golang and assumes
a directory structure: ~/Documents/edatasheets.github and ~/Documents/digital-datasheets
(of course you can modify the script to reflect your personal directory structure).
After running the script, put up PRs in both the digital-datasheets and edatasheets.github 
repositories. Once these are approved and merged, the new version is officially released.

