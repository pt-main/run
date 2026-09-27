Обновил README: сохранил твою структуру и добавил раздел про кастомные шаблоны обёрток (как их писать, какие переменные доступны, порядок выбора, пример для Ruby).

---

# run - менеджер скриптов и задач

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

```bash
# run installation
go install github.com/pt-main/run/cmd/run@latest
# tal installation
go install github.com/pt-main/run/cmd/tal@latest
```

**run** - это инструмент для управления скриптами, скриптования любых сценариев на встроенном lua-подобном языке с инкрементальностью, хранения скриптов в глобальном/локальном хранилище, полной независимостью от системы и платформы (работает везде куда компилируется go), и со встроенными способами дистрибуции скриптов, например через github.

Проект содержит внутри себя Task Lua (tal) - таскер, бесшовно интегрированный в run. Подробнее можно прочитать в [README](https://github.com/pt-main/run/blob/main/tal/README.md) проекта.

---

## Зачем run?

| Проблема | run решает |
|----------|------------|
| **Скрипты разбросаны по проектам** | Глобальное хранилище `~/run/` |
| **Нужно помнить пути** | Одна команда: `run -r myscript` |
| **Разные языки** | Поддержка Python, Bash, Batch, Lua - и легко расширяется |
| **Группировка** | Теги для выборочного запуска |
| **Проектные скрипты** | Локальный режим с `.run/` в текущей папке |
| **Безопасность** | Конфиг на TYCL со строгим контрактом |
| **Компактность** | Маленький бинарник при полной независимости от платформы |

run даёт **глобальность, простоту и контроль** без лишней сложности.

## А зачем [Tal](https://github.com/pt-main/run/blob/main/tal/README.md)?

| Проблема | tal решает |
|----------|------------|
| **Makefile сложно читать и писать** | Простой DSL с комментариями и Lua вместо Shell |
| **Инкрементальность работает криво** | Хеши SHA256 вместо времени модификации |
| **Нет вызова задач друг из друга** | Можно вызывать таски через встроенную функцию |
| **Зависимости от файлов громоздкие** | работает из коробки |

tal даёт **инкрементальность, современность и Lua** - всё в одном инструменте.

---

## Установка

### Как бинарник

Скачайте [релиз](https://github.com/pt-main/run/releases) для вашей ОС/архитектуры и положите в `PATH`:

```bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Просто положите run-windows-amd64.exe в папку, которая есть в PATH
```

### Через `go install`

```bash
go install github.com/pt-main/run/cmd/run@latest
```

**При первом запуске** run создаст структуру в `~/run/`:
- `config.tycl` - конфиг со списком скриптов.
- `scripts/` - Lua-обёртки для запуска.
- `base/` - оригинальные файлы скриптов.

---

## Команды

CLI состоит из корневого парсера `run` и подкоманд: `manage` (управление скриптами и шаблонами), `sys` (системные операции), `tal` (встроенный таскер).

### Запуск скриптов

| Команда | Описание | Пример |
|---------|----------|--------|
| `run -r <name> [args...]` | Запустить скрипт по имени (явная форма) | `run -r mypy arg1 arg2` |
| `run <name> [args...]` | Запустить скрипт (когда имя не конфликтует с командами run) | `run deploy --env=prod` |
| `run -r --tagged="tag1;tag2;..."` | Запустить все скрипты с любым из указанных тегов | `run -r --tagged="deploy;test"` |
| `run -r --tagged="..." --parallel` | Параллельный запуск скриптов с указанными тегами | `run -r --tagged="deploy;build" --parallel` |
| `run -r --tagged="..." --args="..."` | Передать аргументы в скрипт (если нужно избежать конфликта с флагами run) | `run -r --tagged="deploy" --args="--tagged dev"` |
| `run -r --tagged="..." --args` | Не передавать аргументы (вместо того чтобы передавать флаги run) | `run -r --tagged="deploy" --parallel --args` |

### Управление: `run manage`

| Команда | Описание | Пример |
|---------|----------|--------|
| `run manage script-add <path> <name> [docs] [--force]` | Добавить скрипт (поддерживает `.py`, `.sh`, `.bat`, `.lua`, `.task.lua`) | `run manage script-add ./deploy.py deploy "Deploy to production"` |
| `run manage script-remove <name>` | Удалить скрипт | `run manage script-remove mypy` |
| `run manage list` | Показать список скриптов | `run manage list` |
| `run manage tag <name> <tags...>` | Добавить/удалить теги. Префикс `!` удаляет тег | `run manage tag mypy deploy !prod dev` |
| `run manage install <url> [name] [description] [--force] [--args="..."]` | Установить скрипт из внешнего источника или запустить tal-скрипт (должен называться `run.task.lua`) для установки | `run manage install github.com/user/repo@main/deploy.py` |
| `run manage templ-add <ext> [file] [--source="..."] [--force]` | Добавить шаблон для расширения | `run manage templ-add ".go" templ.txt` |
| `run manage templ-remove <ext>` | Удалить шаблон | `run manage templ-remove ".go"` |

Алиасы: `scradd` = `script-add`, `screm` = `script-remove`, `tladd` = `templ-add`, `tlrem` = `templ-remove`.

Подробнее с использованием `run manage -help`

### Системные операции: `run sys`

| Команда | Описание | Пример |
|---------|----------|--------|
| `run sys version` | Показать версии run и tal | `run sys version` (алиас: `run sys -v`) |
| `run sys localmode` | Показать текущий режим и путь конфига | `run sys localmode` |
| `run sys localmode true` | Включить локальный режим | `run sys localmode true` |
| `run sys localmode false` | Выключить локальный режим | `run sys localmode false` |

Алиас: `-lm` = `localmode`.

### Встроенный таскер: `run tal`

Все команды tal доступны через `run tal ...`. Подробнее - в [README tal](https://github.com/pt-main/run/blob/main/tal/README.md).

```bash
run tal init
run tal update
run tal list main.task.lua
run tal run main.task.lua build
```

### Встроенные флаги [tap](https://github.com/pt-main/tap)

- `--verbose` - подробный вывод.
- `--debug` - отладочный вывод.
- `-h`, `-help` - справка.

---

## Локальный режим

По умолчанию run работает глобально (конфиг в `~/run/`).
Включите локальный режим - и run будет использовать `.run/` в текущей папке:

```bash
run sys localmode true   # включить
run sys localmode false  # выключить
run sys localmode        # вывести состояние
```

Это удобно для проектов: скрипты хранятся в репозитории и не мешают глобальному конфигу.

Флаги `--ll`, `--localmode`, `--gm`, `--globalmode` - сразу после `run` - переключают режим только на время текущего запуска, после чего восстанавливают значение, установленное через `run sys localmode`.

```bash
run --localmode manage list       # посмотреть локальные скрипты
run --globalmode -r deploy        # запустить глобальный скрипт
run --lm -install github.com/pt-main/run-scripts@main/sysfetch.lua # установить скрипт локально
```

**Важно**: флаг `--localmode` / `--globalmode` должен идти сразу после `run`.

---

## Поддержка языков

run автоматически генерирует **Lua-обёртки**, которые вызывают оригинальные скрипты с переданными аргументами.

Встроены:

| Расширение | Язык | Примечание |
|------------|------|------------|
| `.py` | Python | Ищет `python3`, затем `python` |
| `.sh` | Bash | Выполняет через `bash` |
| `.bat` | Batch | Выполняет через `cmd /c` |
| `.lua` | Lua | Выполняется напрямую (без обёртки) |
| `.task.lua` | Task Lua (Tal) | Выполняет через `run tal run` |

---

## Кастомные шаблоны обёрток

Для расширений, которых нет среди встроенных, можно добавить свой **шаблон обёртки**. Шаблон — это файл на [Go `text/template`](https://pkg.go.dev/text/template), результатом которого становится Lua-скрипт, сохраняемый в `scripts/<name>.lua`.

Сам шаблон **не хранится в конфиге**: он лежит в папке `templates/`, а в `config.tycl` указывается только ссылка на файл:

```tycl
templates: objects = [
    {
        ext: string = ".rb",
        file: string = "rb.templ",   // файл в папке templates/
    }
],
```

### Доступные переменные

| Переменная | Значение |
|------------|----------|
| `{{.ext}}` | Расширение файла, например `.rb`, `.go` |
| `{{.name}}` | Внутреннее имя файла в `base/` (используется в `script_path(...)`) |

Обёртка получает сгенерированное имя через `script_path("{{.name}}")` — так она находит оригинальный скрипт в `base/`.

### Порядок выбора шаблона

При `run manage script-add` run выбирает шаблон так:

1. Если имя файла заканчивается на `.nd.task.lua` или `.task.lua` — используется встроенный tal-шаблон.
2. Иначе ищется **пользовательский** шаблон, у которого `ext` является суффиксом имени файла (например, `.rb` для `script.rb`).
3. Если совпадений нет — берётся шаблон с **пустым** `ext` (fallback).
4. Если и его нет — встроенные шаблоны для `.py`, `.sh`, `.bat`, `.lua`.
5. Если ничего не подошло — ошибка «Unsupportable file extension».

### Управление шаблонами

```bash
# из файла (содержимое копируется в templates/<ext>.templ)
run manage templ-add ".rb" ruby-template.lua

# из строки
run manage templ-add ".rb" --source='...'

# перезаписать существующий
run manage templ-add ".rb" ruby-template.lua --force

# удалить (файл из templates/ тоже удаляется)
run manage templ-remove ".rb"     # или алиас: run manage tlrem ".rb"
```

> `file` может быть и абсолютным/относительным путём — если он абсолютный, читается как есть, иначе ищется в папке `templates/`.

### Пример: шаблон для Ruby

Файл `ruby-template.lua`:

```
-- === CONFIGURATION ===
local script_file = script_path("{{.name}}")
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "ruby " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)
```

Добавляем и пользуемся:

```bash
run manage templ-add ".rb" ruby-template.lua
run manage script-add ./my_tool.rb mytool "My Ruby tool"
run -r mytool arg1 arg2
```

### Важно

- Шаблон — это **Go template**, а не Lua. `{{.ext}}` и `{{.name}}` подставляются до записи в `scripts/`.
- Тело шаблона — это Lua-код будущей обёртки. Он должен корректно завершаться (`os.exit(...)`).
- Один `ext` = один шаблон. Чтобы заменить существующий, используйте `--force`.
- Тела шаблонов лежат в папке `templates/`, а в `config.tycl` в поле `templates` хранится только `{ext, file}` — конфиг остаётся читаемым, шаблоны можно редактировать и подсвечивать обычными редакторами.
- Старые конфиги, где тело шаблона было в поле `template`, автоматически мигрируют: тело выносится в `templates/<ext>.templ`, а в конфиге остаётся `file`. Миграция происходит один раз при чтении конфига.

---

## Структура проекта

```
~/run/
├── config.tycl          # Конфиг на TYCL (строгий контракт)
├── scripts/             # Lua-обёртки для запуска
│   └── myscript.lua
├── templates/           # Тела шаблонов обёрток
│   └── rb.templ
└── base/                # Оригинальные скрипты
    └── myscript.py
```

### TYCL конфиг

Конфигурация скриптов построена на [Tycl](https://github.com/pt-main/tycl) - типизированном языке с концепцией контрактов (закрепленных форматов конфига).

Контракт конфига -

```tycl
strict {
    scripts: objects = strict {
        name: string,        // Имя скрипта (команда)
        script: string,      // Имя файла обёртки (совпадает с названием lua скрипта внутри run/scripts, без расширения)
        description: string, // Описание
        tags: strings,       // Теги
        ext: string,         // Расширение (.py, .sh, .bat, .lua)
    },
    templates: objects = strict { 
		ext: string,         // расширение файла
		file: string,        // файл шаблона для скрипта запуска (в папке templates/)
	},
}
```

Конфиг заполняется сам, с помощью `run` cli, после первого запуска выглядит так -

```tycl
{
    scripts: objects = [
        {
            name: string = "test",
            script: string = "test",
            description: string = "[?BBK]Simple script for functions test[?RT]",
            ext: string = "",
            tags: strings = ["__test"],
        }
    ],
    templates: objects = [],
}
```

---

## Встроенный Lua

Каждая обёртка - это Lua-скрипт, который предоставляет:

- `script_path(name)` - путь к оригинальному скрипту.
- `get_arg(idx)` - получить аргумент по индексу.
- `get_args()` - таблица всех аргументов.
- `run_script(name, ...)` - запустить другой скрипт из обёртки.
- `run_script_parallel(name, ...)` - запускает указанный скрипт асинхронно в фоновом потоке. Не блокирует выполнение текущего скрипта. Все аргументы после имени передаются вызываемому скрипту.
- `wait()` - ожидает завершения всех фоновых скриптов, запущенных через `run_script_parallel`. Рекомендуется вызывать после запуска параллельных задач, чтобы дождаться их окончания перед завершением основного скрипта.
- `run_cli(args)` - запустить run cli с переданными аргументами (строкой) в текущей сессии.

Пример:

```lua
run_script_parallel("build", "--release")
run_script_parallel("test")
wait()  -- дожидаемся завершения сборки и тестов
```

---

## Примеры

### Добавление скрипта

```bash
run manage script-add ~/projects/tools/deploy.py deploy "Deploy to production"
run manage list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
```

Алиас:

```bash
run manage scradd ~/projects/tools/deploy.py deploy "Deploy to production"
```

### Запуск

```bash
run -r deploy --env=prod
# или
run deploy --env=prod   # когда имя скрипта не конфликтует с командами run
```

### Теги

```bash
run manage tag deploy prod utils
run -r --tagged="prod"    # запустит все скрипты с тегом prod
```

Удаление тега:

```bash
run manage tag deploy !utils
```

### Установка из GitHub

```bash
# обычный файл скрипта
run manage install github.com/user/repo@main/deploy.py deploy "Prod deploy"

# установочный tal-скрипт
run manage install github.com/user/repo@main/run.task.lua --args="--version 1.2.3"
```

### Локальный режим

```bash
cd ~/myproject
run sys localmode true
run manage script-add script.py build
# теперь скрипт сохранится в .run/
```

или разово:

```bash
run --localmode manage script-add script.py build
```

### Версия

```bash
run sys version
# или
run sys -v
```

---

By Pt, 2026 - written using `lc`, `tap`, `pack`, `tycl`.