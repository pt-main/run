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

## Проверка таск-файлов и тегов

```bash
# .task.lua / .nd.task.lua запускаются через run tal
/tmp/runbin manage script-add ./build.task.lua build
/tmp/runbin -r build

# тегированный запуск выполняет все скрипты с тегом
/tmp/runbin manage tag deploy prod
/tmp/runbin manage tag calc   prod
/tmp/runbin -r --tagged="prod"
```

> Если `run` собран против неопубликованной/исправленной версии `tap`,
> подкоманды `manage` и `sys` могут не работать: они не находят свои команды и
> печатают «Has no command». Проверяйте `run manage list` сразу после сборки.
