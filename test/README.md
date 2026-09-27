# Локальный тест `run`

Папка для ручной проверки утилиты в localmode (`~/run` заменяется на `./.run`).

## Запуск

```bash
# из корня репозитория
go build -o /tmp/runbin ./cmd/run

cd test
/tmp/runbin manage list          # .run создастся автоматически при первом запуске
```

> Если `.run` уже существует, автосоздание не происходит — `CheckConfigDir`
> смотрит только на наличие папки. Для чистого прогона удалите её:
> `rm -rf .run`.

## Что проверяем

Шаблоны лежат в `.run/templates/`, в `config.tycl` — только ссылки `{ext, file}`.

```bash
# шаблоны: ruby (.rb) и fallback (пустой ext)
/tmp/runbin manage templ-add ".rb" tpl-ruby.templ
/tmp/runbin manage templ-add ""    tpl-fallback.templ

# скрипты
/tmp/runbin manage script-add ./hello.rb        hello "Ruby hello"
/tmp/runbin manage script-add ./tool.unknownext tool  "Fallback test"
/tmp/runbin manage script-add ./deploy.sh       deploy "Bash test"
/tmp/runbin manage script-add ./calc.py         calc   "Python test"

# запуск
/tmp/runbin -r hello world foo
/tmp/runbin -r tool alpha beta
/tmp/runbin -r deploy prod
```

## Ожидаемое поведение выбора шаблона

| Файл | Шаблон | Команда в обёртке |
|------|--------|-------------------|
| `hello.rb` | `.rb` → `templates/rb.templ` | `ruby <file>` |
| `tool.unknownext` | пустой ext → `templates/default.templ` | `sh <file>` |
| `deploy.sh` | встроенный `.sh` (если fallback удалён) | `bash <file>` |
| `calc.py` | встроенный `.py` (если fallback удалён) | авто-детект `python3`/`python` |

**Важно:** fallback с пустым `ext` проверяется **до** встроенных шаблонов,
поэтому пока он есть в конфиге, `.sh`/`.py`/`.bat` уйдут в него.
Чтобы проверить встроенные — удалите fallback: `manage templ-remove ""`.

## Миграция старого конфига

Если в `config.tycl` тело шаблона лежит в поле `template`, оно автоматически
выносится в `.run/templates/<ext>.templ`, а в конфиге остаётся `file`.
Миграция срабатывает один раз при чтении конфига и идемпотентна.

## Известные особенности (не связаны с хранением шаблонов)

- **Тегированный запуск.** Любая обёртка заканчивается `os.exit(...)`, поэтому
  при `-r --tagged="..."` процесс завершается после **первого** скрипта —
  остальные не выполняются. Это работает только для одиночного запуска.
  С `--parallel` скрипты стартуют, но `os.exit` конкурирует между горутинами
  в общем Lua-состоянии, и вывод может дублироваться.
- **`script-remove`** убирает запись из конфига, но не удаляет файл обёртки
  из `.run/scripts/`.
- **`templ-add` без `--force`** на существующий `ext` возвращает ошибку;
  с `--force` заменяет запись и перезаписывает файл (без дублей).
