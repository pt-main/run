# Files

- [LICENCE](#licence)
- [README-RU.md](#readme-ru-md)
- [README.md](#readme-md)
- [build/run-darwin-amd64-1.0.2](#build-run-darwin-amd64-1-0-2)
- [build/run-darwin-arm64-1.0.2](#build-run-darwin-arm64-1-0-2)
- [build/run-freebsd-amd64-1.0.2](#build-run-freebsd-amd64-1-0-2)
- [build/run-linux-386-1.0.2](#build-run-linux-386-1-0-2)
- [build/run-linux-amd64-1.0.2](#build-run-linux-amd64-1-0-2)
- [build/run-linux-arm-1.0.2](#build-run-linux-arm-1-0-2)
- [build/run-linux-arm64-1.0.2](#build-run-linux-arm64-1-0-2)
- [build/run-openbsd-amd64-1.0.2](#build-run-openbsd-amd64-1-0-2)
- [build/run-windows-386-1.0.2.exe](#build-run-windows-386-1-0-2-exe)
- [build/run-windows-amd64-1.0.2.exe](#build-run-windows-amd64-1-0-2-exe)
- [build.json](#build-json)
- [cmd/run/cli.go](#cmd-run-cli-go)
- [cmd/run/main.go](#cmd-run-main-go)
- [go.mod](#go-mod)
- [go.sum](#go-sum)
- [lib/files.go](#lib-files-go)
- [lib/handlers.go](#lib-handlers-go)
- [lib/localMode/main.go](#lib-localmode-main-go)
- [lib/main.go](#lib-main-go)
- [lib/stdlib.go](#lib-stdlib-go)
- [lib/templates.go](#lib-templates-go)
- [lib/tycl.go](#lib-tycl-go)
- [main.go](#main-go)

---

> Note: all ``` ` ``` symbols was replaced to '

# LICENCE

```
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright [yyyy] [name of copyright owner]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

---

# README-RU.md

```md
# run - менеджер скриптов

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

'''bash
go install github.com/pt-main/run@latest
'''

**run** - это менеджер скриптов, который позволяет добавлять, удалять и запускать скрипты на разных языках одной командой. Скрипты хранятся в '~/run/' и доступны из любой папки.

---

## Зачем run?

| Проблема | run решает |
|----------|------------|
| **Скрипты разбросаны по проектам** | Глобальное хранилище '~/run/' |
| **Нужно помнить пути** | Одна команда: 'run r myscript' |
| **Разные языки** | Поддержка Python, Bash, Batch, Lua - и легко расширяется |
| **Группировка** | Теги для выборочного запуска ('--tagged') |
| **Проектные скрипты** | Локальный режим с '.run/' в текущей папке |
| **Безопасность** | Конфиг на TYCL со строгим контрактом |
| **Компактность** | Маленький бинарник при полной независимости от платформы |

run даёт **глобальность, простоту и контроль** без лишней сложности.

---

## Установка

### Как бинарник

Скачайте [релиз](https://github.com/pt-main/run/releases) для вашей ОС/архитектуры и положите в 'PATH':

'''bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Просто положите run-windows-amd64.exe в папку, которая есть в PATH
'''

### Через 'go install'

'''bash
go install github.com/pt-main/run@latest
'''

**При первом запуске** run создаст структуру в '~/run/':
- 'config.tycl' - конфиг (TYCL) со списком скриптов.
- 'scripts/' - Lua-обёртки для запуска.
- 'base/' - оригинальные файлы скриптов.

---

## Синтаксис

'''bash
run [--<lm/localmode>/<gm/globalmode>] <cmd> <args...>
'''

### Команды

| Команда | Описание | Пример |
|---------|----------|--------|
| 'add <path> <name> [docs]' | Добавить скрипт (поддерживает '.py', '.sh', '.bat', '.lua') | 'run add script.py mypy' |
| 'remove <name>' | Удалить скрипт | 'run remove mypy' |
| 'list' | Показать список скриптов | 'run list' |
| 'r <name> [args...]' | Запустить скрипт | 'run r mypy arg1 arg2' |
| '<name> [args...]' | Запустить скрипт (если имя не совпадает с командой) | 'run mypy arg1' |
| 'tag <name> <tags...>' | Добавить теги | 'run tag mypy deploy prod' |
| 'rm-tag <name> <tags...>' | Удалить теги | 'run rm-tag mypy prod' |
| 'localmode [true/false]' | Включить/выключить локальный режим, показать текущее состояние запуска скриптов | 'run localmode true' |
| 'r --tagged="tag1;tag2;..."' | Запустить скрипты с любым из тегов | 'run r --tagged="deploy;test"' |
| 'r --tagged="..." --parralel' | Параллельный запуск скриптов с нужным тегом | 'run r --tagged="deploy;build" --paralel' |
| 'r --tagged="..." --args=""' | Передача аргументов в скрипт (если нужно избежать конфликта, например с флагами run, или не передавать аргументы) | 'run r --tagged="deploy;build" --args="--tagged dev"','run r --tagged="deploy;build" --parallel --args' - не передает аргументов, вместо того чтобы передавать '--parallel' |
| 'version' | Показать версию | 'run version' |

### Флаги

- '--force' при 'add' - заменяет существующий скрипт с таким же именем.
- '--tagged="tag1;tag2"' при 'r' - запуск по тегам.
- '--ll / --localmode / --gm / --globalmode' сразу после 'run' - запуск в локальном/глобальном режиме, после завершения востанавливает установленный с помощью 'run localmode' режим.

---

## Локальный режим

По умолчанию run работает глобально (конфиг в '~/run/').  
Включите локальный режим - и run будет использовать '.run/' в текущей папке:

'''bash
run localmode true  # включить
run localmode false # выключить
run localmode       # fasle - выводит состояние
'''

Это удобно для проектов: скрипты хранятся в репозитории и не мешают глобальному конфигу.

---

## Поддержка языков

run автоматически генерирует **Lua-обёртки**, которые вызывают оригинальные скрипты с переданными аргументами.

| Расширение | Язык | Примечание |
|------------|------|------------|
| '.py' | Python | Ищет 'python3', затем 'python' |
| '.sh' | Bash | Выполняет через 'bash' |
| '.bat' | Batch (Windows) | Выполняет через 'cmd /c' |
| '.lua' | Lua | Выполняется напрямую (без обёртки) |

Добавить новый язык просто - достаточно дописать шаблон в 'templates.go'. 

---

## Структура проекта

'''
~/run/
├── config.tycl          # Конфиг на TYCL (строгий контракт)
├── scripts/             # Lua-обёртки для запуска
│   └── myscript.lua
└── base/                # Оригинальные скрипты
    └── myscript.py
'''


### TYCL конфиг

Конфигурация скриптов построена на [Tycl](https://github.com/pt-main/tycl) - типизированном языке с концепйией контрактов (закрепленных форматов конфига). 

Контракт конфига - 

'''tycl
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
'''

Конфиг заполняется сам, с помощью 'run' cli, после первого запуска выглядит так - 

'''tycl
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
'''

---

## Встроенный Lua

Каждая обёртка - это Lua-скрипт, который предоставляет:

- 'script_path(name)' - путь к оригинальному скрипту.
- 'get_arg(idx)' - получить аргумент по индексу.
- 'get_args()' - таблица всех аргументов.
- 'run_script(name, ...)' - запустить другой скрипт из обёртки.

---

## Примеры

### Добавление скрипта

'''bash
run add ~/projects/tools/deploy.py deploy "Deploy to production"
run list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
'''

### Запуск

'''bash
run r deploy --env=prod
# или
run deploy --env=prod # когда имя скрипта не конфликтует с командами run
'''

### Теги

'''bash
run tag deploy prod utils
run r --tagged="prod"   # запустит все скрипты с тегом prod
'''

### Локальный режим

'''bash
cd ~/myproject
run localmode true
run add script.py build
# теперь скрипт сохранится в .run/
'''

или 

'''bash
run --localmode add script.py build
'''

**Важно**: флаг '--localmode' для корректной должен быть сразу после 'run'.

---

## Лицензия

Apache 2.0 - подробности в [LICENSE](LICENSE).

---

By Pt, 2026 - написано с использованием tap, tycl и lc.
```

---

# README.md

```md
# run – Script Manager

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

'''bash
go install github.com/pt-main/run@latest
'''

**run** is a script manager that lets you add, remove, and run scripts in various languages with a single command. Scripts are stored in '~/run/' and are accessible from any folder.

---

## Why run?

| Problem | run solves |
|---------|------------|
| **Scripts scattered across projects** | Global storage at '~/run/' |
| **Need to remember paths** | One command: 'run r myscript' |
| **Different languages** | Supports Python, Bash, Batch, Lua – and easily extensible |
| **Grouping** | Tags for selective execution ('--tagged') |
| **Project scripts** | Local mode with '.run/' in the current folder |
| **Safety** | TYCL config with a strict contract |
| **Compactness** | Small binary, fully platform-independent |

run gives **globality, simplicity, and control** without unnecessary complexity.

---

## Installation

### As a binary

Download the [release](https://github.com/pt-main/run/releases) for your OS/architecture and place it in your 'PATH':

'''bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Just put run-windows-amd64.exe in a folder that is in your PATH
'''

### Via 'go install'

'''bash
go install github.com/pt-main/run@latest
'''

**On first launch** run will create the following structure in '~/run/':
- 'config.tycl' – configuration (TYCL) with the list of scripts.
- 'scripts/' – Lua wrappers for execution.
- 'base/' – original script files.

---

## Syntax

'''bash
run [--<lm>/<localmode> | --<gm>/<globalmode>] <cmd> <args...>
'''

### Commands

| Command | Description | Example |
|---------|-------------|---------|
| 'add <path> <name> [docs]' | Add a script (supports '.py', '.sh', '.bat', '.lua') | 'run add script.py mypy' |
| 'remove <name>' | Remove a script | 'run remove mypy' |
| 'list' | List all scripts | 'run list' |
| 'r <name> [args...]' | Run a script | 'run r mypy arg1 arg2' |
| '<name> [args...]' | Run a script (if the name does not conflict with a run command) | 'run mypy arg1' |
| 'tag <name> <tags...>' | Add tags | 'run tag mypy deploy prod' |
| 'rm-tag <name> <tags...>' | Remove tags | 'run rm-tag mypy prod' |
| 'localmode [true/false]' | Enable/disable local mode, show current state | 'run localmode true' |
| 'r --tagged="tag1;tag2;..."' | Run scripts that have any of the given tags | 'run r --tagged="deploy;test"' |
| 'r --tagged="..." --parallel' | Run tagged scripts in parallel | 'run r --tagged="deploy;build" --parallel' |
| 'r --tagged="..." --args=""' | Pass arguments to the script (use when you need to avoid conflicts with run flags, or to explicitly pass no arguments) | 'run r --tagged="deploy;build" --args="--tagged dev"' <br> 'run r --tagged="deploy;build" --parallel --args' – passes no arguments (instead of passing '--parallel') |
| 'version' | Show version | 'run version' |

### Flags

- '--force' with 'add' – overwrite an existing script with the same name.
- '--tagged="tag1;tag2"' with 'r' – run by tags.
- '--ll' / '--localmode' / '--gm' / '--globalmode' immediately after 'run' – run in local/global mode for that single command, restoring the previously set mode afterwards.

---

## Local mode

By default, run works globally (config in '~/run/').  
Enable local mode and run will use '.run/' in the current folder:

'''bash
run localmode true  # enable
run localmode false # disable
run localmode       # prints current state (e.g., false)
'''

This is convenient for projects: scripts live in the repository and do not interfere with the global config.

---

## Language support

run automatically generates **Lua wrappers** that invoke the original scripts with the passed arguments.

| Extension | Language | Notes |
|-----------|----------|-------|
| '.py'     | Python   | Looks for 'python3', then 'python' |
| '.sh'     | Bash     | Executes via 'bash' |
| '.bat'    | Batch (Windows) | Executes via 'cmd /c' |
| '.lua'    | Lua      | Executed directly (no wrapper) |

Adding a new language is easy – just add a template in 'templates.go'.

---

## Project structure

'''
~/run/
├── config.tycl          # TYCL config (strict contract)
├── scripts/             # Lua wrappers for execution
│   └── myscript.lua
└── base/                # Original scripts
    └── myscript.py
'''

### TYCL config

Script configuration is built on [Tycl](https://github.com/pt-main/tycl) – a typed language with the concept of contracts (fixed config formats).

The config contract is:

'''tycl
strict {
    scripts: objects = strict {
        name: string,        // Script name (command)
        script: string,      // Wrapper file name (matches the Lua script name in run/scripts, without extension)
        description: string, // Description
        tags: strings,       // Tags
        source: string,      // Source path of the original script
        ext: string,         // Extension (.py, .sh, .bat, .lua)
    },
}
'''

The config is filled automatically by the 'run' CLI. After first launch it looks like this:

'''tycl
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
'''

---

## Built‑in Lua

Each wrapper is a Lua script that provides:

- 'script_path(name)' – path to the original script.
- 'get_arg(idx)' – get an argument by index.
- 'get_args()' – table of all arguments.
- 'run_script(name, ...)' – run another script from the wrapper.

---

## Examples

### Adding a script

'''bash
run add ~/projects/tools/deploy.py deploy "Deploy to production"
run list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
'''

### Running

'''bash
run r deploy --env=prod
# or
run deploy --env=prod   # when the script name does not conflict with run commands
'''

### Tags

'''bash
run tag deploy prod utils
run r --tagged="prod"   # runs all scripts with the tag prod
'''

### Local mode

'''bash
cd ~/myproject
run localmode true
run add script.py build
# now the script will be saved in .run/
'''

or

'''bash
run --localmode add script.py build
'''

**Important:** the '--localmode' flag must appear immediately after 'run' to work correctly.

---

## License

Apache 2.0 – see [LICENSE](LICENSE) for details.

---

By Pt, 2026 – built with tap, tycl and lc.
```

---

# build/run-darwin-amd64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xcf in position 0: invalid continuation byte
```

---

# build/run-darwin-arm64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xcf in position 0: invalid continuation byte
```

---

# build/run-freebsd-amd64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xc8 in position 40: invalid continuation byte
```

---

# build/run-linux-386-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xc0 in position 24: invalid start byte
```

---

# build/run-linux-amd64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0x90 in position 40: invalid start byte
```

---

# build/run-linux-arm-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0x9c in position 24: invalid start byte
```

---

# build/run-linux-arm64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xb7 in position 18: invalid start byte
```

---

# build/run-openbsd-amd64-1.0.2

```2
Can't read file: 'utf-8' codec can't decode byte 0xc0 in position 24: invalid start byte
```

---

# build/run-windows-386-1.0.2.exe

```exe
Can't read file: 'utf-8' codec can't decode byte 0x90 in position 2: invalid start byte
```

---

# build/run-windows-amd64-1.0.2.exe

```exe
Can't read file: 'utf-8' codec can't decode byte 0x90 in position 2: invalid start byte
```

---

# build.json

```json
{
  "output_dir": "./build",
  "platforms": "all",
  "version": "1.0.2",
  "cgo_enabled": 0,
  "tags": "netgo,osusergo",
  "go_build_args": ["-trimpath", "-buildvcs=false"],
  "ldflags": "-s -w",
  "verbose": false
}
```

---

# cmd/run/cli.go

```go
package main

import (
	"fmt"

	"github.com/pt-main/run"
	runlib "github.com/pt-main/run/lib"
	localmode "github.com/pt-main/run/lib/localMode"
	"github.com/pt-main/tap"
	"github.com/pt-main/tycl/cli"
)

func NewCli() *tap.Parser {
	p := tap.NewParser("run", '[?BE]╭─────── [?BRD]Run[?RT]
[?BE]⎬─ [?RT]Simple and powerfull script manager
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]

[?GN]Usage: [?RT]
  [?BYW]Type: [?BBK]run [--<lm|locamode>] <cmd> <args...>

  [?BBK]run <cmd> <args...> [?YW]- execute run command (which described in help)[?RT]
    [?BYW]Example: [?BBK]run -add script.py script1[?RT]
  
  [?BBK]run r <script> <args...> [?YW]- run registred script[?RT]
    [?BYW]Example: [?BBK]run r script1 --os='linux'[?RT]

  [?BBK]run <script> <args...> [?YW]- run registred script which name isn't same with cli commands[?RT]
    [?BYW]Example: [?BBK]run r script1 --os='linux'[?RT]
  
  [?BBK]run --tagged="<tag1>;<tag2>;<...>" <args...>[?YW] - run script with any tag[?RT]
    [?BYW]Example: [?BBK]run r --tagged="deploy" --os='linux'[?RT]', []string{"-h", "help"}, tap.DefaultParserConfig())

	p.AddCommand("-tycl", func(p *tap.Parser, s []string) error {
		args, err := runlib.ProcessShell(s[0])
		if err != nil {
			return err
		}
		if err = cli.NewCli().Parse(args); err != nil {
			return err
		}
		return nil
	}, tap.DONT_SHOW, []string{"code"}, nil, false)

	p.AddCommand("add", runlib.AddHandler,
		'Add script (python/bash) with auto run script generation. 
Use --force flag to replace script if it's already added with same name.',
		[]string{"path", "name"}, []string{"docs"}, false)

	p.AddCommand("remove", runlib.RemoveHandler,
		'Remove script from global config',
		[]string{"name"}, nil, false)

	p.AddCommand("list", runlib.ListHandler,
		'Show scripts list', nil, nil, false)

	p.AddCommand("r",
		runlib.MakeRunHandler(true), 'Run scripts',
		nil, nil, true)

	p.AddCommand(tap.DEFAULT_CMD,
		runlib.MakeRunHandler(false), '',
		nil, nil, true)

	p.AddCommand("tag",
		runlib.TagHahdler, 'Tag script',
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("rm-tag",
		runlib.RmTagHahdler, 'Remove tags from script',
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("version", func(p *tap.Parser, s []string) error {
		fmt.Println("run v" + run.Version + ", by Pt, Apache 2.0 licence")
		return nil
	}, 'Show version and info', nil, nil, false)

	p.AddCommand("localmode", func(p *tap.Parser, s []string) error {
		if len(s) == 0 {
			fmt.Println("localmode:", localmode.IsLocalmode(), "| path:", runlib.ConfigDirPath())
			return nil
		}
		switch s[0] {
		case "true":
			localmode.Set(true)
		case "false":
			localmode.Set(false)
		default:
			return fmt.Errorf("Invalid argument")
		}
		return nil
	}, 'Set run working mode. Mode - 'true'/'false'', nil, []string{"mode"}, false)

	return p
}
```

---

# cmd/run/main.go

```go
package main

import (
	"fmt"
	"os"

	run "github.com/pt-main/run/lib"
	localmode "github.com/pt-main/run/lib/localMode"
)

func main() {
	args := os.Args
	lm := localmode.IsLocalmode()
	temp := lm

	if len(args) > 1 {
		if args[1] == "--localmode" || args[1] == "--lm" {
			temp = true
		} else if args[1] == "--globalmode" || args[1] == "--gm" {
			temp = false
		}
	}
	localmode.Set(temp)

	ok, err := run.CheckConfigDir()
	if err != nil {
		fmt.Println("Can't check installation:", err)
		return
	}
	if !ok {
		if err := run.InstallConfigDir(); err != nil {
			fmt.Println("Can't make run dir:", err)
			return
		}
	}

	err = NewCli().Main()
	if err != nil {
		fmt.Println(err)
	}

	if localmode.IsLocalmode() == temp {
		localmode.Set(lm)
	}
}
```

---

# go.mod

```mod
module github.com/pt-main/run

go 1.24.13

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/dlclark/regexp2 v1.12.0 // indirect
	github.com/mattn/go-shellwords v1.0.14 // indirect
	github.com/pt-main/lc v1.5.4 // indirect
	github.com/pt-main/tap v1.4.10 // indirect
	github.com/pt-main/tycl v1.3.7 // indirect
	github.com/yuin/gopher-lua v1.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

---

# go.sum

```sum
github.com/BurntSushi/toml v1.6.0 h1:dRaEfpa2VI55EwlIW72hMRHdWouJeRF7TPYhI+AUQjk=
github.com/BurntSushi/toml v1.6.0/go.mod h1:ukJfTF/6rtPPRCnwkur4qwRxa8vTRFBF0uk2lLoLwho=
github.com/dlclark/regexp2 v1.12.0 h1:0j4c5qQmnC6XOWNjP3PIXURXN2gWx76rd3KvgdPkCz8=
github.com/dlclark/regexp2 v1.12.0/go.mod h1:DHkYz0B9wPfa6wondMfaivmHpzrQ3v9q8cnmRbL6yW8=
github.com/mattn/go-shellwords v1.0.14 h1:yUKzIgsCnosndOASY6/enly1EAuaXeFSQ7cdyA3OuYg=
github.com/mattn/go-shellwords v1.0.14/go.mod h1:EZzvwXDESEeg03EKmM+RmDnNOPKG4lLtQsUlTZDWQ8Y=
github.com/pt-main/lc v1.5.4 h1:PLrlu7uvKP4lcNB/xxPghnY91vxxf7X1PJB0vNQskk0=
github.com/pt-main/lc v1.5.4/go.mod h1:uUxWI4oiOkia6Tko+cgF+O3fhGZF3of1NSHeTEjOaMU=
github.com/pt-main/tap v1.4.8 h1:BBT83yIiE08qZNvT32gMfUAjT7We+OjNpuegXUDEjno=
github.com/pt-main/tap v1.4.8/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.4.10 h1:4ROxXKlNAYV6MBjbjZfKM4Dyu2m3RUuTj3/AAYMFKwk=
github.com/pt-main/tap v1.4.10/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tycl v1.3.5 h1:QZkZnL1Aa/IHHADVMXaN5sZR/sF4OtuOvSdm8DV9TZI=
github.com/pt-main/tycl v1.3.5/go.mod h1:IIu0EBQkmHRH2b2w2KpYdY3uLfEIFqq9jioOpafJYnU=
github.com/pt-main/tycl v1.3.6 h1:ozHyCzwOTCcb2eZpPVx/AFvq96hwc2WgCfCkeks21to=
github.com/pt-main/tycl v1.3.6/go.mod h1:IIu0EBQkmHRH2b2w2KpYdY3uLfEIFqq9jioOpafJYnU=
github.com/pt-main/tycl v1.3.7 h1:atjOrwIUkoMvCpmOivE+QRvkZq2ePcDuY0uKynmXr1k=
github.com/pt-main/tycl v1.3.7/go.mod h1:IIu0EBQkmHRH2b2w2KpYdY3uLfEIFqq9jioOpafJYnU=
github.com/yuin/gopher-lua v1.1.2 h1:yF/FjE3hD65tBbt0VXLE13HWS9h34fdzJmrWRXwobGA=
github.com/yuin/gopher-lua v1.1.2/go.mod h1:7aRmXIWl37SqRf0koeyylBEzJ+aPt8A+mmkQ4f1ntR8=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

---

# lib/files.go

```go
package runlib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-shellwords"
	localmode "github.com/pt-main/run/lib/localMode"
	"github.com/pt-main/tycl/generation"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func ConfigDirPath() string {
	p, err := os.UserHomeDir()
	dir := "run"
	if localmode.IsLocalmode() {
		p, err = os.Getwd()
		dir = ".run"
	}
	if err != nil {
		panic(err)
	}
	return filepath.Join(p, dir)
}

func ConfigDirScriptsPath() string {
	return filepath.Join(ConfigDirPath(), "scripts")
}

func ConfigDirBasePath() string {
	return filepath.Join(ConfigDirPath(), "base")
}

func ConfigDirConfigPath() string {
	return filepath.Join(ConfigDirPath(), "config.tycl")
}

func ProcessShell(cmdStr string) ([]string, error) {
	args, err := shellwords.Parse(cmdStr)
	if err != nil {
		return nil, fmt.Errorf("Parse shell args: %v", err)
	}
	if len(args) == 0 {
		return nil, nil
	}
	return args, nil
}

func CheckConfigDir() (bool, error) {
	info, err := os.Stat(ConfigDirPath())
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func NewScriptConfig(name, script, description, source, ext string, tags []string) *shared.Config {
	conf := shared.NewNilConfig()
	conf.StringV["name"] = name
	conf.StringV["script"] = script
	conf.StringV["description"] = description
	conf.StringV["source"] = source
	conf.StringV["ext"] = ext
	if tags == nil {
		tags = []string{}
	}
	conf.StringArrV["tags"] = tags
	return conf
}

func FormatConfig(config *shared.Config) (res string, err error) {
	if _, ok := config.InnerArrV["scripts"]; !ok {
		config.InnerArrV["scripts"] = make([]*shared.Config, 0)
	}
	res, err = generation.Tycl(config)
	return
}

func NewRunScript(name, content string) error {
	return utils.WriteF(filepath.Join(ConfigDirScriptsPath(), name+".lua"), content)
}

func NewScript(name, content string) error {
	return utils.WriteF(filepath.Join(ConfigDirBasePath(), name), content)
}

func UpdateConfig(config *shared.Config) error {
	conf, err := FormatConfig(config)
	if err != nil {
		return err
	}
	if err := utils.WriteF(ConfigDirConfigPath(), conf); err != nil {
		return err
	}
	return nil
}

func InstallConfigDir() error {
	if err := os.Mkdir(ConfigDirPath(), 0755); err != nil {
		return err
	}
	if err := os.Mkdir(ConfigDirScriptsPath(), 0755); err != nil {
		return err
	}
	if err := os.Mkdir(ConfigDirBasePath(), 0755); err != nil {
		return err
	}
	conf, err := StdLib()
	if err != nil {
		return err
	}
	if err := utils.WriteF(ConfigDirConfigPath(), conf); err != nil {
		return err
	}
	return nil
}
```

---

# lib/handlers.go

```go
package runlib

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/pt-main/tap"
	"github.com/pt-main/tap/color"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func AddHandler(p *tap.Parser, s []string) error {
	ext := filepath.Ext(s[0])
	fpsplit := strings.Split(s[0], "/")
	file := fpsplit[0]
	if len(fpsplit) > 0 {
		file = fpsplit[len(fpsplit)-1]
	}
	runScript := ""
	conf, err := GetCfg()
	if err != nil {
		return err
	}
	_, force := p.Flags["force"]
	newScripts := []*shared.Config{}
	for _, script := range conf.InnerArrV["scripts"] {
		name := script.StringV["script"]
		if name == s[1] && !force {
			return fmt.Errorf("Can't add script: script already added. Use --force to replace script.")
		}
		if name != s[1] {
			newScripts = append(newScripts, script)
		}
	}
	conf.InnerArrV["scripts"] = newScripts
	script, err := utils.OpenF(s[0])
	if err != nil {
		return err
	}
	addScript := true
	switch ext {
	case ".py":
		runScript = PythonRunScriptTemplate(s[0])
	case ".sh":
		runScript = BashRunScriptTemplate(s[0])
	case ".bat":
		runScript = BatRunScriptTemplate(s[1])
	case ".lua":
		runScript = script
		addScript = false
	default:
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}
	docs := ""
	if len(s) > 2 {
		docs = s[2]
	}
	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"], NewScriptConfig(s[1], s[1], docs, script, ext, nil))
	if err := NewRunScript(s[1], runScript); err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	if addScript {
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
	return nil
}

func RemoveHandler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	newScripts := []*shared.Config{}
	for _, script := range cfg.InnerArrV["scripts"] {
		name := script.StringV["name"]
		if name != s[0] {
			newScripts = append(newScripts, script)
		}
	}
	cfg.InnerArrV["scripts"] = newScripts
	return UpdateConfig(cfg)
}

func ListHandler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	color.PrintlnColored("[?GN]╭─────── [?YW] Scripts [?RT]")
	linestart := "[?GN]│     [?RT]"
	for _, script := range cfg.InnerArrV["scripts"] {
		name := script.StringV["name"]
		ext := script.StringV["ext"]
		if ext != "" {
			ext = "[?BBK] (" + ext + ")"
		}
		description := script.StringV["description"]
		color.PrintColored("[?GN]⎬─ [?YW]%v%v[?RT]", name, ext)
		if description != "" {
			color.PrintlnColored(":\n"+linestart+"%v[?RT]", strings.ReplaceAll(description, "\n", "\n"+linestart))
		} else {
			fmt.Println()
		}
	}
	color.PrintlnColored("[?GN]╰───────[?RT]")
	return nil
}

func RunScript(cfg *shared.Config, name string, rArgs []string) error {
	var scriptPath string
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		scriptPath_ := script.StringV["script"]
		if scriptName == name {
			scriptPath = scriptPath_
			break
		}
	}
	if scriptPath == "" {
		return fmt.Errorf("Script is not found")
	}
	file, err := utils.OpenF(filepath.Join(ConfigDirScriptsPath(), scriptPath+".lua"))
	if err != nil {
		return err
	}
	if err := NewLuaState(rArgs).DoString(file); err != nil {
		return err
	}
	return nil
}

func MakeRunHandler(hasR bool) func(p *tap.Parser, s []string) error {
	return func(p *tap.Parser, s []string) error {
		cfg, err := GetCfg()
		if err != nil {
			return err
		}
		idx := 1
		if slices.Contains([]string{"--gm", "--globalmode",
			"--lm", "--localmode"}, p.RawArgs[0]) {
			idx += 1
		}
		if hasR {
			idx += 1
		}
		var args []string = nil
		args_, ok := p.Flags["args"]
		if ok {
			args, err = ProcessShell(args_)
			if err != nil {
				return err
			}
			if args == nil {
				args = []string{}
			}
		}
		if args == nil {
			args = p.RawArgs[idx:]
		}
		if tags_, ok := p.Flags["tagged"]; ok {
			_, parallel := p.Flags["parallel"]
			tags := strings.Split(tags_, ";")
			errs := []string{}
			goru := 0
			for _, script := range cfg.InnerArrV["scripts"] {
				scrTags := script.StringArrV["tags"]
				scriptName := script.StringV["name"]
				for _, tag := range scrTags {
					if slices.Contains(tags, tag) {
						p.Print("verbose", "Run %v: ", scriptName)
						if parallel {
							go func() {
								goru += 1
								if err := RunScript(cfg, scriptName, args); err != nil {
									p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
									errs = append(errs, err.Error())
								} else {
									p.Print("verbose", "[?GN]Ok[?RT]")
								}
								goru -= 1
							}()
							time.Sleep(time.Second / 100)
						} else {
							if err := RunScript(cfg, scriptName, args); err != nil {
								p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
								errs = append(errs, err.Error())
							} else {
								p.Print("verbose", "[?GN]Ok[?RT]")
							}
						}
					}
				}
			}
			p.Print("verbose", "[?GN]Gorutines: %v[?RT]", goru)
			for goru != 0 {
				time.Sleep(time.Second / 100)
			}
			if len(errs) == 0 {
				return nil
			}
			return fmt.Errorf(" - " + strings.Join(errs, "\n - "))
		} else {
			if len(s) < 1 {
				return fmt.Errorf("Invalid argument length: need more or equals to 1")
			}
			name := s[0]
			p.Print("verbose", "Run %v: ", name)
			return RunScript(cfg, name, args)
		}
	}
}

func RmTagHahdler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			newTags := []string{}
			for _, tag := range script.StringArrV["tags"] {
				for _, rm := range s[1:] {
					if tag != rm {
						newTags = append(newTags, tag)
					}
				}
			}
			script.StringArrV["tags"] = newTags
			break
		}
	}
	return UpdateConfig(cfg)
}

func TagHahdler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			script.StringArrV["tags"] = append(script.StringArrV["tags"], s[1:]...)
			break
		}
	}
	return UpdateConfig(cfg)
}
```

---

# lib/localMode/main.go

```go
package localmode

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pt-main/tycl/utils"
)

func ConfigLocalmodePath() string {
	p, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(p, "run.localmode")
}

func CheckConfigLocalmode() bool {
	_, err := os.Stat(ConfigLocalmodePath())
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func Install() {
	if !CheckConfigLocalmode() {
		utils.WriteF(ConfigLocalmodePath(), "false")
	}
}

func IsLocalmode() bool {
	Install()
	file, err := utils.OpenF(ConfigLocalmodePath())
	if err != nil {
		return false
	}
	return strings.TrimSpace(file) == "true"
}

func Set(local bool) {
	Install()
	content := "false"
	if local {
		content = "true"
	}
	if err := utils.WriteF(ConfigLocalmodePath(), content); err != nil {
		panic(err)
	}
}
```

---

# lib/main.go

```go
package runlib

import (
	"fmt"
	"path/filepath"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
	lua "github.com/yuin/gopher-lua"
)

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	var errI core.ErrorInterface
	cfg, errI := tycl.Process(file, tyclContract, true)
	if errI != nil {
		return cfg, fmt.Errorf(format.FormatError(errI))
	}
	return cfg, nil
}

func NewLuaState(args []string) *lua.LState {
	L := lua.NewState()

	L.SetGlobal("script_path", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		path := filepath.Join(ConfigDirBasePath(), name)
		L.Push(lua.LString(path))
		return 1
	}))

	L.SetGlobal("get_arg", L.NewFunction(func(L *lua.LState) int {
		idx := L.CheckInt(1)
		if idx < 1 || idx > len(args) {
			L.Push(lua.LNil)
		} else {
			L.Push(lua.LString(args[idx-1]))
		}
		return 1
	}))

	L.SetGlobal("get_args", L.NewFunction(func(L *lua.LState) int {
		tbl := L.NewTable()
		for i, arg := range args {
			tbl.RawSetInt(i+1, lua.LString(arg))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("run_script", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		var scriptArgs []string
		top := L.GetTop()
		for i := 2; i <= top; i++ {
			arg := L.Get(i)
			if str, ok := arg.(lua.LString); ok {
				scriptArgs = append(scriptArgs, string(str))
			} else {
				scriptArgs = append(scriptArgs, L.ToStringMeta(arg).String())
			}
		}

		cfg, err := GetCfg()
		if err != nil {
			L.RaiseError("failed to load config: %v", err)
			return 0
		}

		if err := RunScript(cfg, name, scriptArgs); err != nil {
			L.RaiseError("failed to run script %q: %v", name, err)
			return 0
		}
		return 0
	}))

	return L
}
```

---

# lib/stdlib.go

```go
package runlib

import (
	"github.com/pt-main/tycl/shared"
)

func StdLib() (string, error) {
	config := shared.NewNilConfig()
	NewRunScript("test", 'print("test script"); print(script_path("test.py")); print(get_args()[1])')
	config.InnerArrV["scripts"] = append(config.InnerArrV["scripts"], NewScriptConfig(
		"test", "test", "[?BBK]Simple script for functions test[?RT]", "", "", []string{"__test"},
	))
	conf, err := FormatConfig(config)
	if err != nil {
		return "", err
	}
	return conf, nil
}
```

---

# lib/templates.go

```go
package runlib

import "fmt"

func PythonRunScriptTemplate(name string) string {
	return fmt.Sprintf('-- === CONFIGURATION ===
local script_file = script_path(%v)
local args = get_args()
-- =====================

local function get_python()
    local function check(cmd)
        local f = io.popen(cmd .. " --version 2>&1")
        if f then
            local out = f:read("*a")
            f:close()
            return out:match("Python") ~= nil
        end
        return false
    end
    if check("python3") then return "python3" end
    if check("python")   then return "python"  end
    return nil
end

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local python = get_python()
if not python then
    io.stderr:write("Error: Python interpreter not found\n")
    os.exit(1)
end

local cmd = python .. " " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)', fmt.Sprintf("%#v", name))
}

func BashRunScriptTemplate(name string) string {
	return fmt.Sprintf('-- === CONFIGURATION ===
local script_file = script_path(%v)
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "bash " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)', fmt.Sprintf("%#v", name))
}

func BatRunScriptTemplate(name string) string {
	return fmt.Sprintf('-- === CONFIGURATION ===
local script_file = script_path(%v)
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "cmd /c " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)', fmt.Sprintf("%#v", name))
}

func LuaRunScriptTemplate(name string) string {
	return fmt.Sprintf('-- === CONFIGURATION ===
local script_file = script_path(%v)
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "lua " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)', fmt.Sprintf("%#v", name))
}
```

---

# lib/tycl.go

```go
package runlib

var tyclContract = '
strict {
	scripts: objects = strict {
		name: string,
		script: string,
		description: string,
		tags: strings,
		source: string,
		ext: string,
	},
}
'
```

---

# main.go

```go
package run

var Version = "1.1.0"
```

---

