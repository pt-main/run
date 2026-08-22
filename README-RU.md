# run - менеджер скриптов и задач

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

```bash
# run installation
go install github.com/pt-main/run/cmd/run@latest
# tal installation
go install github.com/pt-main/run/tal/cmd/tal@latest
```

**run** - это менеджер скриптов, который позволяет добавлять, удалять и запускать скрипты на разных языках одной командой. Скрипты хранятся в `~/run/` и доступны из любой папки.

Проект содержит внутри себя Task Lua (tal) - таскер, безшовно интегрированный в run. Подробнее можно прочитать в [README](https://github.com/pt-main/run/blob/main/tal/README.md) проекта.

---

## Зачем run?

| Проблема | run решает |
|----------|------------|
| **Скрипты разбросаны по проектам** | Глобальное хранилище `~/run/` |
| **Нужно помнить пути** | Одна команда: `run -r myscript` |
| **Разные языки** | Поддержка Python, Bash, Batch, Lua - и легко расширяется |
| **Группировка** | Теги для выборочного запуска (`--tagged`) |
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
go install github.com/pt-main/run@latest
```

**При первом запуске** run создаст структуру в `~/run/`:
- `config.tycl` - конфиг (TYCL) со списком скриптов.
- `scripts/` - Lua-обёртки для запуска.
- `base/` - оригинальные файлы скриптов.

---

## Синтаксис

```bash
run [--<lm/localmode>/<gm/globalmode>] <cmd> <args...>
```

### Команды

| Команда | Описание | Пример |
|---------|----------|--------|
| `-add <path> <name> [docs]` | Добавить скрипт (поддерживает `.py`, `.sh`, `.bat`, `.lua`) | `run -add script.py mypy` |
| `-remove <name>` | Удалить скрипт | `run -remove mypy` |
| `-list` | Показать список скриптов | `run -list` |
| `-r <name> [args...]` | Запустить скрипт | `run -r mypy arg1 arg2` |
| `<name> [args...]` | Запустить скрипт (если имя не совпадает с командой) | `run mypy arg1` |
| `-tag <name> <tags...>` | Добавить теги | `run -tag mypy deploy prod` |
| `-rm-tag <name> <tags...>` | Удалить теги | `run -rm-tag mypy prod` |
| `-localmode [true/false]` | Включить/выключить локальный режим, показать текущее состояние запуска скриптов | `run -localmode true` |
| `-r --tagged="tag1;tag2;..."` | Запустить скрипты с любым из тегов | `run -r --tagged="deploy;test"` |
| `-r --tagged="..." --parralel` | Параллельный запуск скриптов с нужным тегом | `run -r --tagged="deploy;build" --paralel` |
| `-r --tagged="..." --args=""` | Передача аргументов в скрипт (если нужно избежать конфликта, например с флагами run, или не передавать аргументы) | `run -r --tagged="deploy;build" --args="--tagged dev"`,`run -r --tagged="deploy;build" --parallel --args` - не передает аргументов, вместо того чтобы передавать `--parallel` |
| `-version` | Показать версию | `run -version` |

### Флаги

- `--force` при `-add` - заменяет существующий скрипт с таким же именем.
- `--tagged="tag1;tag2"` при `-r` - запуск по тегам.
- `--ll / --localmode / --gm / --globalmode` сразу после `run` - запуск в локальном/глобальном режиме, после завершения востанавливает установленный с помощью `run -localmode` режим.

---

## Локальный режим

По умолчанию run работает глобально (конфиг в `~/run/`).  
Включите локальный режим - и run будет использовать `.run/` в текущей папке:

```bash
run -localmode true  # включить
run -localmode false # выключить
run -localmode       # fasle - выводит состояние
```

Это удобно для проектов: скрипты хранятся в репозитории и не мешают глобальному конфигу.

---

## Поддержка языков

run автоматически генерирует **Lua-обёртки**, которые вызывают оригинальные скрипты с переданными аргументами.

| Расширение | Язык | Примечание |
|------------|------|------------|
| `.py` | Python | Ищет `python3`, затем `python` |
| `.sh` | Bash | Выполняет через `bash` |
| `.bat` | Batch | Выполняет через `cmd /c` |
| `.lua` | Lua | Выполняется напрямую (без обёртки) |
| `.tal.lua` | Task Lua (Tal) | Выполняет через `run tal` |

---

## Структура проекта

```
~/run/
├── config.tycl          # Конфиг на TYCL (строгий контракт)
├── scripts/             # Lua-обёртки для запуска
│   └── myscript.lua
└── base/                # Оригинальные скрипты
    └── myscript.py
```


### TYCL конфиг

Конфигурация скриптов построена на [Tycl](https://github.com/pt-main/tycl) - типизированном языке с концепйией контрактов (закрепленных форматов конфига). 

Контракт конфига - 

```tycl
strict {
    scripts: objects = strict {
        name: string,        // Имя скрипта (команда)
        script: string,      // Имя файла обёртки (совпадает с названием lua скрипта внутри run/scripts, без расширения)
        description: string, // Описание
        tags: strings,       // Теги
        source: string,      // Исходник оригинального скрипта
        ext: string,         // Расширение (.py, .sh, .bat, .lua)
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
            source: string = "",
            ext: string = "",
            tags: strings = ["__test"],
        }
    ],
}
```

---

## Встроенный Lua

Каждая обёртка - это Lua-скрипт, который предоставляет:

- `script_path(name)` - путь к оригинальному скрипту.
- `get_arg(idx)` - получить аргумент по индексу.
- `get_args()` - таблица всех аргументов.
- `run_script(name, ...)` - запустить другой скрипт из обёртки.

---

## Примеры

### Добавление скрипта

```bash
run -add ~/projects/tools/deploy.py deploy "Deploy to production"
run -list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
```

### Запуск

```bash
run -r deploy --env=prod
# или
run deploy --env=prod # когда имя скрипта не конфликтует с командами run
```

### Теги

```bash
run -tag deploy prod utils
run -r --tagged="prod"   # запустит все скрипты с тегом prod
```

### Локальный режим

```bash
cd ~/myproject
run -localmode true
run -add script.py build
# теперь скрипт сохранится в .run/
```

или 

```bash
run --localmode add script.py build
```

**Важно**: флаг `--localmode` для корректной должен быть сразу после `-run`.

---

By Pt, 2026 – written using `lc`, `tap`, `pack`, `tycl`.
