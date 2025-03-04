# Gital

A git repo manager for the developer who is constantly switching between projects.

## Why?

I am a platform engineer for a company with a good number of microservices and terraform repos. This means that I am always switching between different git repos to make small infrastructure and devops based changes. I wanted to create something that allows me to easily search and open repos without having to navigate to them using a file explorer. Hence Gital.

This is not a Git Client like [GitKraken](https://www.gitkraken.com/) or [RilaGit](https://rela.dev/).

## The Idea

Three components will be used to achieve the functionality I want, first is a background service that will scan a designated folder (maybe more than one) for git repos periodically. It will monitor and store information on all these repos in a database, so it can be searched through. Next is a 'FatUI', a way to view the repos, perform certain bulk tasks (like bulk git fetch) and configure ways of opening that repo. It's called a FatUI because it doesn't need to be super performant, and therefore can be created using cross-platform technology like Electron. Lastly is the 'finder', a spotlight type UI that allows you to quickly fuzzy search through your repos and 'open' them. This UI needs to be performant and lightweight. It's the UI that'll be used 90% of the time, initialised on start up so it's warm when needed, and is activated on a global hotkey. Because of the performance requirements, this will most likely be a native app, and therefor needs a bespoke app per operating system (starting with Windows). 

## Features

- [ ] Automatically scan for new repos
- [ ] Set custom 'open' commands per repo (with sensible default)
- [ ] Ability to add non-git folders
- [ ] Fuzzy search through git repos using spotlight-style search widget
- [ ] Useful bulk repo commands like fetch all, and set all to `main`
- [ ] Ability to export and import all repos, making onboarding and new laptop setups easier
- [ ] Cross-platform
  - [ ] Windows
  - [ ] Linux
  - [ ] IOS

