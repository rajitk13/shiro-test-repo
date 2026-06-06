# Shiro Test Repository

Test repository for validating Shiro's AI-powered code review and GitHub inline commenting features.

## Purpose

This minimal Go project demonstrates automated code review workflows using Gemini AI provider integrated with GitHub Actions.

## Tech Stack

- Go
- Shiro Automation
- Gemini AI (Google AI Studio)
- GitHub Actions

## Setup

This repository uses Shiro for automated code review on pull requests. When a PR is created, Shiro will:
1. Get the diff of changes
2. Send it to Gemini AI for review
3. Post inline comments on specific lines with feedback
