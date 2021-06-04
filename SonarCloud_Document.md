# Sonarcloud Document

## SonarCloud

Besides all the various linters for Go, TypeScript, and other languages included in system builds and in VSCode via plugins,
we also have dedicated static analysis tools.

SonarCloud is a cloud-based code analysis service designed to detect code quality issues in different programming languages, 
continuously ensuring the maintainability, reliability and security of your code.

Early detection of problems ensures that fewer issues get through to the later stages of the process and ultimately helps to increase the overall quality of your production code.

----
## Setting up Sonarcloud

### Set up your organization

Connect your GitHub organization with SonarCloud
When prompted, install the SonarCloud application on GitHub. This step allows SonarCloud to access your GitHub organization or personal account.
You can select specific repositories to be connected to SonarCloud or just select all. You can always change this settings later.

In this step, you will set up the SonarCloud organization that corresponds to your GitHub organization or personal account.

SonarCloud will suggest a key for your SonarCloud organization. This is a name unique across all organizations within SonarCloud. 
Add this Sonar key to your Github Repository which you want to anayze.
Go to Settings -> Secret-> Add Sonar token.

we have configured sonarcloud with automate with GitHub Action method.
First, you need to sign in into sonarcloud with our existing GitHub account on the repository service that hosts the code you want to analyze. 
Once you have successfully logged in you will see the SonarCloud welcome screen. Click on Import projects from GitHub.

On the SonarCloud Dashboard, Click on + option at the right hand side to select  project which you want to analyze.

----
## Configuring Analysis with SonarCloud
 
Recommend to Configure analysis with GitHub Action method and follow the steps recommended in SonarCloud Platform.
 
After follwing all the steps, Analysis will start on the dasboard in few minutes.

After any merge in the master, SonarCloud will run the analysis. It will analyze master branch and the open PR'S.
 
 ----
## Excluding Files from Code Analysis
 
For exclusing some of the files to not get scanned during analysis, we have to set file exclusion property in 
Sonar-project.properties file.

These are the files which we have already excluded:
 ```
 sonar.exclusions=**/*.pb.*.go, **/*.pb.go, **/test*/**, .gitignore, .git/**, .semgrep/**, **/*.bindata.go, **/*.spec.ts, dev-docs/**, .buildkite/*, .expeditor/**, .studio/*, benchmarks/*, bin/*, cache/*, dev/*, ec2/*, protovendor/*, results/*, scripts/*, terraform/*, tools/*, **/*.sql, **/*.docs-chef-io, **/*.md
```
 
## Coverage Report generation.
 
To generate code coverage report we have to set Coverage report property in  Sonar-project.properties file.
 
We have set it for go and javascript files:
``` 
sonar.go.coverage.reportPaths=Coverage_report/cover.out
sonar.javascript.lcov.reportPaths=Coverage_report/lcov.info
```

We have to keep the report files in Coverage_report folder at root directory and set the path accordingly in Sonar-project.properties files.

https://docs.sonarqube.org/latest/analysis/coverage/

 
Code coverage report will be shown in SonarCloud dashboard which will be a single report for both the files. 
  
  
 



