# multi-file-comparer

Usecase: If working on multiple project and some files across projects must be the same.

Config file format:

```
rootdir: /home/user/src
subfolders: [project1, project2, project3]
files:
  - .github/actions/linters/.eslintrc.yaml
```

Above config make sure that below files are the same, otherwise will write warning to the console.

- /home/user/src/project1/.github/actions/linters/.eslintrc.yaml
- /home/user/src/project2/.github/actions/linters/.eslintrc.yaml
- /home/user/src/project3/.github/actions/linters/.eslintrc.yaml

