## ライブラリ
### next
- SWR

### go
- gin
- gorm

## run
### build image
```bash
docker compose build
```

### run container
```bash
docker compose run
```

### run container in background

If you want to keep using the terminal, you can run containers in detached mode.

```bash
docker compose run -d
```

### stop container

If you run containers in detached mode, you can stop containers like this:

```bash
docker compose down
```

### fetch yarn.lock

If `yarn.lock` does not exist, execute the following command:

```bash
docker run --rm unitimetable-uni-nextjs cat /app/yarn.lock > frontend/yarn.lock
```
