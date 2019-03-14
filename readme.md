Commit
======

Find something to collaborate

run
---

1.  Set working environment variable:
    `commit` project using three different working environment:

- `development`
- `test`
- `production`
  to set working environment, export `COMMIT_ENV` variable. Default value is `development`.

2.  Edit config file of your working environment.

3.  Build Dockerfile. You can pass working environment variable with `--build-arg app_env=development`. Use your working environment instead of development.
