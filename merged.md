# Files

- [LICENSE](#license)
- [README-RU.md](#readme-ru-md)
- [README.md](#readme-md)
- [build.json](#build-json)
- [cmd/run/main.go](#cmd-run-main-go)
- [cmd/tal/main.go](#cmd-tal-main-go)
- [go.mod](#go-mod)
- [go.sum](#go-sum)
- [main.go](#main-go)
- [rubytempl.txt](#rubytempl-txt)
- [run/api/funcs.go](#run-api-funcs-go)
- [run/api/localMode/main.go](#run-api-localmode-main-go)
- [run/api/lua.go](#run-api-lua-go)
- [run/api/main.go](#run-api-main-go)
- [run/api/stdlib.go](#run-api-stdlib-go)
- [run/api/templates.go](#run-api-templates-go)
- [run/api/tycl.go](#run-api-tycl-go)
- [run/runcli/cli.go](#run-runcli-cli-go)
- [run/runcli/handlers/handlers.go](#run-runcli-handlers-handlers-go)
- [run/runcli/handlers/install.go](#run-runcli-handlers-install-go)
- [run/runcli/handlers/templates.go](#run-runcli-handlers-templates-go)
- [run/runcli/manage.go](#run-runcli-manage-go)
- [run/runcli/process.go](#run-runcli-process-go)
- [run/runcli/sys.go](#run-runcli-sys-go)
- [tal/README-ru.md](#tal-readme-ru-md)
- [tal/README.md](#tal-readme-md)
- [tal/core/main.go](#tal-core-main-go)
- [tal/generation/generate.go](#tal-generation-generate-go)
- [tal/generation/runtime.go](#tal-generation-runtime-go)
- [tal/lang/errFmt.go](#tal-lang-errfmt-go)
- [tal/lang/lcproc.go](#tal-lang-lcproc-go)
- [tal/lang/process.go](#tal-lang-process-go)
- [tal/lang/struct.go](#tal-lang-struct-go)
- [tal/lua/main.go](#tal-lua-main-go)
- [tal/main.go](#tal-main-go)
- [tal/runtime/main.go](#tal-runtime-main-go)
- [tal/shared/config.go](#tal-shared-config-go)
- [tal/shared/errors.go](#tal-shared-errors-go)
- [tal/test/main.go](#tal-test-main-go)
- [tal/test/main.task.lua](#tal-test-main-task-lua)
- [tal/test/simple.task.lua](#tal-test-simple-task-lua)

---

> Note: all ``` ` ``` symbols was replaced to '

# LICENSE

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
Обновил README: сохранил твою структуру и добавил раздел про кастомные шаблоны обёрток (как их писать, какие переменные доступны, порядок выбора, пример для Ruby).

---

# run - менеджер скриптов и задач

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

'''bash
# run installation
go install github.com/pt-main/run/cmd/run@latest
# tal installation
go install github.com/pt-main/run/cmd/tal@latest
'''

**run** - это инструмент для управления скриптами, скриптования любых сценариев на встроенном lua-подобном языке с инкрементальностью, хранения скриптов в глобальном/локальном хранилище, полной независимостью от системы и платформы (работает везде куда компилируется go), и со встроенными способами дистрибуции скриптов, например через github.

Проект содержит внутри себя Task Lua (tal) - таскер, бесшовно интегрированный в run. Подробнее можно прочитать в [README](https://github.com/pt-main/run/blob/main/tal/README.md) проекта.

---

## Зачем run?

| Проблема | run решает |
|----------|------------|
| **Скрипты разбросаны по проектам** | Глобальное хранилище '~/run/' |
| **Нужно помнить пути** | Одна команда: 'run -r myscript' |
| **Разные языки** | Поддержка Python, Bash, Batch, Lua - и легко расширяется |
| **Группировка** | Теги для выборочного запуска |
| **Проектные скрипты** | Локальный режим с '.run/' в текущей папке |
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
go install github.com/pt-main/run/cmd/run@latest
'''

**При первом запуске** run создаст структуру в '~/run/':
- 'config.tycl' - конфиг со списком скриптов.
- 'scripts/' - Lua-обёртки для запуска.
- 'base/' - оригинальные файлы скриптов.

---

## Команды

CLI состоит из корневого парсера 'run' и подкоманд: 'manage' (управление скриптами и шаблонами), 'sys' (системные операции), 'tal' (встроенный таскер).

### Запуск скриптов

| Команда | Описание | Пример |
|---------|----------|--------|
| 'run -r <name> [args...]' | Запустить скрипт по имени (явная форма) | 'run -r mypy arg1 arg2' |
| 'run <name> [args...]' | Запустить скрипт (когда имя не конфликтует с командами run) | 'run deploy --env=prod' |
| 'run -r --tagged="tag1;tag2;..."' | Запустить все скрипты с любым из указанных тегов | 'run -r --tagged="deploy;test"' |
| 'run -r --tagged="..." --parallel' | Параллельный запуск скриптов с указанными тегами | 'run -r --tagged="deploy;build" --parallel' |
| 'run -r --tagged="..." --args="..."' | Передать аргументы в скрипт (если нужно избежать конфликта с флагами run) | 'run -r --tagged="deploy" --args="--tagged dev"' |
| 'run -r --tagged="..." --args' | Не передавать аргументы (вместо того чтобы передавать флаги run) | 'run -r --tagged="deploy" --parallel --args' |

### Управление: 'run manage'

| Команда | Описание | Пример |
|---------|----------|--------|
| 'run manage script-add <path> <name> [docs] [--force]' | Добавить скрипт (поддерживает '.py', '.sh', '.bat', '.lua', '.task.lua') | 'run manage script-add ./deploy.py deploy "Deploy to production"' |
| 'run manage script-remove <name>' | Удалить скрипт | 'run manage script-remove mypy' |
| 'run manage list' | Показать список скриптов | 'run manage list' |
| 'run manage tag <name> <tags...>' | Добавить/удалить теги. Префикс '!' удаляет тег | 'run manage tag mypy deploy !prod dev' |
| 'run manage install <url> [name] [description] [--force] [--args="..."]' | Установить скрипт из внешнего источника или запустить tal-скрипт (должен называться 'run.task.lua') для установки | 'run manage install github.com/user/repo@main/deploy.py' |
| 'run manage templ-add <ext> [file] [--source="..."] [--force]' | Добавить шаблон для расширения | 'run manage templ-add ".go" templ.txt' |
| 'run manage templ-remove <ext>' | Удалить шаблон | 'run manage templ-remove ".go"' |

Алиасы: 'scradd' = 'script-add', 'screm' = 'script-remove', 'tladd' = 'templ-add', 'tlrem' = 'templ-remove'.

Подробнее с использованием 'run manage -help'

### Системные операции: 'run sys'

| Команда | Описание | Пример |
|---------|----------|--------|
| 'run sys version' | Показать версии run и tal | 'run sys version' (алиас: 'run sys -v') |
| 'run sys localmode' | Показать текущий режим и путь конфига | 'run sys localmode' |
| 'run sys localmode true' | Включить локальный режим | 'run sys localmode true' |
| 'run sys localmode false' | Выключить локальный режим | 'run sys localmode false' |

Алиас: '-lm' = 'localmode'.

### Встроенный таскер: 'run tal'

Все команды tal доступны через 'run tal ...'. Подробнее - в [README tal](https://github.com/pt-main/run/blob/main/tal/README.md).

'''bash
run tal init
run tal update
run tal list main.task.lua
run tal run main.task.lua build
'''

### Встроенные флаги [tap](https://github.com/pt-main/tap)

- '--verbose' - подробный вывод.
- '--debug' - отладочный вывод.
- '-h', '-help' - справка.

---

## Локальный режим

По умолчанию run работает глобально (конфиг в '~/run/').
Включите локальный режим - и run будет использовать '.run/' в текущей папке:

'''bash
run sys localmode true   # включить
run sys localmode false  # выключить
run sys localmode        # вывести состояние
'''

Это удобно для проектов: скрипты хранятся в репозитории и не мешают глобальному конфигу.

Флаги '--ll', '--localmode', '--gm', '--globalmode' - сразу после 'run' - переключают режим только на время текущего запуска, после чего восстанавливают значение, установленное через 'run sys localmode'.

'''bash
run --localmode manage list       # посмотреть локальные скрипты
run --globalmode -r deploy        # запустить глобальный скрипт
run --lm -install github.com/pt-main/run-scripts@main/sysfetch.lua # установить скрипт локально
'''

**Важно**: флаг '--localmode' / '--globalmode' должен идти сразу после 'run'.

---

## Поддержка языков

run автоматически генерирует **Lua-обёртки**, которые вызывают оригинальные скрипты с переданными аргументами.

Встроены:

| Расширение | Язык | Примечание |
|------------|------|------------|
| '.py' | Python | Ищет 'python3', затем 'python' |
| '.sh' | Bash | Выполняет через 'bash' |
| '.bat' | Batch | Выполняет через 'cmd /c' |
| '.lua' | Lua | Выполняется напрямую (без обёртки) |
| '.task.lua' | Task Lua (Tal) | Выполняет через 'run tal run' |

---

## Кастомные шаблоны обёрток

Для расширений, которых нет среди встроенных, можно добавить свой **шаблон обёртки**. Шаблон — это обычная строка на [Go 'text/template'](https://pkg.go.dev/text/template), результатом которой становится Lua-скрипт, сохраняемый в 'scripts/<name>.lua'.

### Доступные переменные

| Переменная | Значение |
|------------|----------|
| '{{.ext}}' | Расширение файла, например '.rb', '.go' |
| '{{.name}}' | Внутреннее имя файла в 'base/' (используется в 'script_path(...)') |

Обёртка получает сгенерированное имя через 'script_path("{{.name}}")' — так она находит оригинальный скрипт в 'base/'.

### Порядок выбора шаблона

При 'run manage script-add' run выбирает шаблон так:

1. Если имя файла заканчивается на '.nd.task.lua' или '.task.lua' — используется встроенный tal-шаблон.
2. Иначе ищется **пользовательский** шаблон, у которого 'ext' является суффиксом имени файла (например, '.rb' для 'script.rb').
3. Если совпадений нет — берётся шаблон с **пустым** 'ext' (fallback).
4. Если и его нет — встроенные шаблоны для '.py', '.sh', '.bat', '.lua'.
5. Если ничего не подошло — ошибка «Unsupportable file extension».

### Управление шаблонами

'''bash
# из файла
run manage templ-add ".rb" ruby-template.lua

# из строки
run manage templ-add ".rb" --source='...'

# перезаписать существующий
run manage templ-add ".rb" ruby-template.lua --force

# удалить
run manage templ-remove ".rb"     # или алиас: run manage tlrem ".rb"
'''

### Пример: шаблон для Ruby

Файл 'ruby-template.lua':

'''
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
'''

Добавляем и пользуемся:

'''bash
run manage templ-add ".rb" ruby-template.lua
run manage script-add ./my_tool.rb mytool "My Ruby tool"
run -r mytool arg1 arg2
'''

### Fallback-шаблон

Шаблон с пустым 'ext' вызывается, когда ни один другой не совпал. Удобно для универсальной обёртки:

'''bash
run manage templ-add "" --source='-- fallback wrapper
local f = script_path("{{.name}}")
local args = get_args()
local cmd = f
for _, a in ipairs(args) do
    cmd = cmd .. " " .. '"' .. a .. '"'
end
os.exit(os.execute(cmd) or 0)'
'''

### Важно

- Шаблон — это **Go template**, а не Lua. '{{.ext}}' и '{{.name}}' подставляются до записи в 'scripts/'.
- Тело шаблона — это Lua-код будущей обёртки. Он должен корректно завершаться ('os.exit(...)').
- Один 'ext' = один шаблон. Чтобы заменить существующий, используйте '--force'.
- Шаблоны хранятся в 'config.tycl' в поле 'templates' и не зависят от языка оригинала.

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

Конфигурация скриптов построена на [Tycl](https://github.com/pt-main/tycl) - типизированном языке с концепцией контрактов (закрепленных форматов конфига).

Контракт конфига -

'''tycl
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
		template: string,    // шаблон для скрипта запуска
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
            ext: string = "",
            tags: strings = ["__test"],
        }
    ],
    templates: objects = [],
}
'''

---

## Встроенный Lua

Каждая обёртка - это Lua-скрипт, который предоставляет:

- 'script_path(name)' - путь к оригинальному скрипту.
- 'get_arg(idx)' - получить аргумент по индексу.
- 'get_args()' - таблица всех аргументов.
- 'run_script(name, ...)' - запустить другой скрипт из обёртки.
- 'run_script_parallel(name, ...)' - запускает указанный скрипт асинхронно в фоновом потоке. Не блокирует выполнение текущего скрипта. Все аргументы после имени передаются вызываемому скрипту.
- 'wait()' - ожидает завершения всех фоновых скриптов, запущенных через 'run_script_parallel'. Рекомендуется вызывать после запуска параллельных задач, чтобы дождаться их окончания перед завершением основного скрипта.
- 'run_cli(args)' - запустить run cli с переданными аргументами (строкой) в текущей сессии.

Пример:

'''lua
run_script_parallel("build", "--release")
run_script_parallel("test")
wait()  -- дожидаемся завершения сборки и тестов
'''

---

## Примеры

### Добавление скрипта

'''bash
run manage script-add ~/projects/tools/deploy.py deploy "Deploy to production"
run manage list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
'''

Алиас:

'''bash
run manage scradd ~/projects/tools/deploy.py deploy "Deploy to production"
'''

### Запуск

'''bash
run -r deploy --env=prod
# или
run deploy --env=prod   # когда имя скрипта не конфликтует с командами run
'''

### Теги

'''bash
run manage tag deploy prod utils
run -r --tagged="prod"    # запустит все скрипты с тегом prod
'''

Удаление тега:

'''bash
run manage tag deploy !utils
'''

### Установка из GitHub

'''bash
# обычный файл скрипта
run manage install github.com/user/repo@main/deploy.py deploy "Prod deploy"

# установочный tal-скрипт
run manage install github.com/user/repo@main/run.task.lua --args="--version 1.2.3"
'''

### Локальный режим

'''bash
cd ~/myproject
run sys localmode true
run manage script-add script.py build
# теперь скрипт сохранится в .run/
'''

или разово:

'''bash
run --localmode manage script-add script.py build
'''

### Версия

'''bash
run sys version
# или
run sys -v
'''

---

By Pt, 2026 - written using 'lc', 'tap', 'pack', 'tycl'.
```

---

# README.md

```md
# run - script and task manager

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

'''bash
# run installation
go install github.com/pt-main/run/cmd/run@latest
# tal installation
go install github.com/pt-main/run/cmd/tal@latest
'''

**run** is a tool for managing scripts, scripting any scenarios in an embedded Lua-like language with incrementality, storing scripts in global/local storage, complete independence from system and platform (works anywhere Go compiles), and with built-in ways to distribute scripts, for example via GitHub.

The project contains Task Lua (tal) inside itself - a task runner seamlessly integrated into run. More details can be read in the project [README](https://github.com/pt-main/run/blob/main/tal/README.md).

---

## Why run?

| Problem | run solves |
|----------|------------|
| **Scripts scattered across projects** | Global storage '~/run/' |
| **Need to remember paths** | One command: 'run -r myscript' |
| **Different languages** | Support for Python, Bash, Batch, Lua - and easily extensible |
| **Grouping** | Tags for selective running ('--tagged') |
| **Project scripts** | Local mode with '.run/' in the current folder |
| **Security** | TYCL config with a strict contract |
| **Compactness** | Small binary while fully platform-independent |

run gives **globality, simplicity and control** without unnecessary complexity.

## And why [Tal](https://github.com/pt-main/run/blob/main/tal/README.md)?

| Problem | tal solves |
|----------|------------|
| **Makefile is hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incrementality works poorly** | SHA256 hashes instead of modification time |
| **No calling tasks from each other** | Tasks can be called via a built-in function |
| **File dependencies are cumbersome** | works out of the box |

tal gives **incrementality, modernity and Lua** - all in one tool.

---

## Installation

### As a binary

Download the [release](https://github.com/pt-main/run/releases) for your OS/architecture and put it in 'PATH':

'''bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Just put run-windows-amd64.exe in a folder that is in PATH
'''

### Via 'go install'

'''bash
go install github.com/pt-main/run@latest
'''

**On first launch** run will create a structure in '~/run/':
- 'config.tycl' - config with the list of scripts.
- 'scripts/' - Lua wrappers for launching.
- 'base/' - original script files.

---


## Commands

| Command | Description | Example |
|---------|-------------|---------|
| '-add <path> <name> [docs] [--force]' | Add a script (supports '.py', '.sh', '.bat', '.lua') | 'run -add script.py mypy' |
| '-remove <name>' | Remove a script | 'run -remove mypy' |
| '-list' | Show the list of scripts | 'run -list' |
| '-install <url> [name] [description] [--force] [--args="..."]' | Install a script from an external source, or run a tal script for installation | |
| '<name> [args...]' | Run a script (if the name does not match a command) | 'run mypy arg1' |
| '-tag <name> <tags...>' | Add/remove tags. Use the '!' prefix for a tag to remove it. | 'run -tag mypy deploy prod' |
| '-localmode [true/false]' | Enable/disable local mode, show the current script launch state | 'run -localmode true' |
| '-r <name> [args...] [--tagged='...']' | Run a script | 'run -r mypy arg1 arg2' |
| '-r --tagged="tag1;tag2;..."' | Run scripts with any of the tags | 'run -r --tagged="deploy;test"' |
| '-r --tagged="..." --parallel' | Run scripts with the required tag in parallel | 'run -r --tagged="deploy;build" --parallel' |
| '-r --tagged="..." --args=""' | Pass arguments to the script (if you need to avoid a conflict, for example with run flags, or not pass arguments) | 'run -r --tagged="deploy;build" --args="--tagged dev"','run -r --tagged="deploy;build" --parallel --args' - does not pass arguments instead of passing '--parallel' |
| '-version' | Show the version of run and tal | 'run -version' |

'--no_color' - flag disables colored output throughout the session.

---

## Local mode

By default run works globally (config in '~/run/').  
Enable local mode - and run will use '.run/' in the current folder:

'''bash
run -localmode true  # enable
run -localmode false # disable
run -localmode       # show state
'''

This is convenient for projects: scripts are stored in the repository and do not interfere with the global config.

'--ll / --localmode / --gm / --globalmode' immediately after 'run' - launch in local/global mode; after completion, restores the mode set with 'run -localmode'.

---

## Language support

run automatically generates **Lua wrappers** that call the original scripts with the passed arguments.

| Extension | Language | Note |
|------------|------|------------|
| '.py' | Python | Looks for 'python3', then 'python' |
| '.sh' | Bash | Executes via 'bash' |
| '.bat' | Batch | Executes via 'cmd /c' |
| '.lua' | Lua | Executes directly (without a wrapper) |
| '.task.lua' | Task Lua (Tal) | Executes via 'run tal run' |

---

## Project structure

'''
~/run/
├── config.tycl          # Config in TYCL (strict contract)
├── scripts/             # Lua wrappers for launching
│   └── myscript.lua
└── base/                # Original scripts
    └── myscript.py
'''


### TYCL config

Script configuration is built on [Tycl](https://github.com/pt-main/tycl) - a typed language with the concept of contracts (fixed config formats).

Config contract -

'''tycl
strict {
    scripts: objects = strict {
        name: string,        // Script name (command)
        script: string,      // Name of the wrapper file (matches the Lua script name inside run/scripts, without extension)
        description: string, // Description
        tags: strings,       // Tags
        ext: string,         // Extension (.py, .sh, .bat, .lua)
    },
}
'''

The config is filled in automatically by the 'run' CLI; after the first launch it looks like this -

'''tycl
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
}
'''

---

## Built-in Lua

Each wrapper is a Lua script that provides:

- 'script_path(name)' - path to the original script.
- 'get_arg(idx)' - get an argument by index.
- 'get_args()' - table of all arguments.
- 'run_script(name, ...)' - run another script from the wrapper.
- 'run_script_parallel(name, ...)' - runs the specified script asynchronously in a background thread. Does not block execution of the current script. All arguments after the name are passed to the called script.
- 'wait()' - waits for all background scripts started via 'run_script_parallel' to finish. It is recommended to call it after starting parallel tasks to wait for their completion before the main script exits.
- 'run_cli(args)' - run run cli with the passed arguments (as a string) in the current session.

Example:
'''lua
run_script_parallel("build", "--release")
run_script_parallel("test")
wait()  -- wait for the build and tests to finish
'''

---

## Examples

### Adding a script

'''bash
run -add ~/projects/tools/deploy.py deploy "Deploy to production"
run -list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
'''

### Running

'''bash
run -r deploy --env=prod
# or
run deploy --env=prod # when the script name does not conflict with run commands
'''

### Tags

'''bash
run -tag deploy prod utils
run -r --tagged="prod"   # will run all scripts with the prod tag
'''

### Local mode

'''bash
cd ~/myproject
run -localmode true
run -add script.py build
# now the script will be saved in .run/
'''

or

'''bash
run --localmode add script.py build
'''

**Important**: for correct operation, the '--localmode' flag must be immediately after 'run'.

---

By Pt, 2026 - written using 'lc', 'tap', 'pack', 'tycl'.
```

---

# build.json

```json
{
  "project_path": "./cmd/run",
  "output_dir": "./build",
  "name_template": "{project}-v{version}-{os}-{arch}",
  "version": "1.3.4",
  "clean_before_build": true,
  "cgo_enabled": 0,
  "go_build_args": ["-trimpath", "-buildvcs=false", "-gcflags=all=-l"],
  "ldflags": "-s -w -X main.version=${version} -X main.commit=${commit} -X main.buildTime=${date}T${time}",
  "tags": "",
  "platforms": "all",
  "exclude_platforms": ["android/*", "ios/*", "plan9/*"],
  "platform_config": {
    "linux/arm": {
      "goarm": "7",
      "ldflags": "-s -w -X main.arch=armv7"
    },
    "linux/amd64": {
      "goamd64": "v3",
      "ldflags": "-s -w -X main.arch=amd64v3"
    },
    "windows/amd64": {
      "cgo_enabled": 1,
      "tags": "windows",
      "env": {
        "CC": "x86_64-w64-mingw32-gcc"
      }
    },
    "darwin/amd64": {
      "goamd64": "v2",
      "env": {
        "CGO_ENABLED": "1",
        "SDKROOT": "/path/to/macOS.sdk"
      }
    }
  },
  "verbose": true
}
```

---

# cmd/run/main.go

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pt-main/run/run/runcli"
)

func main() {
	cli, err := runcli.NewCli()
	if err != nil {
		log.Fatal("SYSTEM ERROR: CREATING CLI:\n", err)
		return
	}
	err = runcli.Process(cli, os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
```

---

# cmd/tal/main.go

```go
package main

import (
	"fmt"

	"github.com/pt-main/run/tal/runtime"
)

func main() {
	p := runtime.CreateCli()
	if err := p.Main(); err != nil {
		fmt.Println(err)
	}
}
```

---

# go.mod

```mod
module github.com/pt-main/run

go 1.24.13

require (
	github.com/bmatcuk/doublestar/v4 v4.10.0
	github.com/dlclark/regexp2 v1.12.0
	github.com/iancoleman/orderedmap v0.3.0
	github.com/mattn/go-shellwords v1.0.14
	github.com/pt-main/lc v1.5.7-f
	github.com/pt-main/pack v1.2.0
	github.com/pt-main/tap v1.4.14-0.20260922135441-c0b7d2c122cf
	github.com/pt-main/tycl v1.3.8
	github.com/yuin/gopher-lua v1.1.2
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/pt-main/tap/go v1.5.8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

---

# go.sum

```sum
github.com/BurntSushi/toml v1.6.0 h1:dRaEfpa2VI55EwlIW72hMRHdWouJeRF7TPYhI+AUQjk=
github.com/BurntSushi/toml v1.6.0/go.mod h1:ukJfTF/6rtPPRCnwkur4qwRxa8vTRFBF0uk2lLoLwho=
github.com/bmatcuk/doublestar/v4 v4.10.0 h1:zU9WiOla1YA122oLM6i4EXvGW62DvKZVxIe6TYWexEs=
github.com/bmatcuk/doublestar/v4 v4.10.0/go.mod h1:xBQ8jztBU6kakFMg+8WGxn0c6z1fTSPVIjEY1Wr7jzc=
github.com/dlclark/regexp2 v1.12.0 h1:0j4c5qQmnC6XOWNjP3PIXURXN2gWx76rd3KvgdPkCz8=
github.com/dlclark/regexp2 v1.12.0/go.mod h1:DHkYz0B9wPfa6wondMfaivmHpzrQ3v9q8cnmRbL6yW8=
github.com/iancoleman/orderedmap v0.3.0 h1:5cbR2grmZR/DiVt+VJopEhtVs9YGInGIxAoMJn+Ichc=
github.com/iancoleman/orderedmap v0.3.0/go.mod h1:XuLcCUkdL5owUCQeF2Ue9uuw1EptkJDkXXS7VoV7XGE=
github.com/mattn/go-shellwords v1.0.14 h1:yUKzIgsCnosndOASY6/enly1EAuaXeFSQ7cdyA3OuYg=
github.com/mattn/go-shellwords v1.0.14/go.mod h1:EZzvwXDESEeg03EKmM+RmDnNOPKG4lLtQsUlTZDWQ8Y=
github.com/pt-main/lc v1.5.6 h1:YLRawKqIIQX4b3hFiPHfgTk7jvv1PRAk8RofWfBG8ZA=
github.com/pt-main/lc v1.5.6/go.mod h1:uUxWI4oiOkia6Tko+cgF+O3fhGZF3of1NSHeTEjOaMU=
github.com/pt-main/lc v1.5.7-f h1:YD9nMphPA62n7ztrgxvA7nTOloDi7rXJPuKL0AD/L84=
github.com/pt-main/lc v1.5.7-f/go.mod h1:uUxWI4oiOkia6Tko+cgF+O3fhGZF3of1NSHeTEjOaMU=
github.com/pt-main/pack v1.1.2 h1:I5mHCd3ax1dR7c7DeOvStpf6tuVNmA4iXgV9RIDfFWM=
github.com/pt-main/pack v1.1.2/go.mod h1:RDJ+eUeINksFnliDhZGGDTxxbTB5aZoQW9nu4Ndz8UY=
github.com/pt-main/pack v1.1.5 h1:v/hmUd/jy8cgf3e7/2PcBKE/v0J7hPplTFoeiCN7Tpc=
github.com/pt-main/pack v1.1.5/go.mod h1:kvGNuw7AZOPGjCujs9gc/CkiYXoC78atKT1sI/hvqFA=
github.com/pt-main/pack v1.1.6 h1:3r5ywoVfGA1dUpfb5OvnTrgergN79p81uowru3b85u8=
github.com/pt-main/pack v1.1.6/go.mod h1:kvGNuw7AZOPGjCujs9gc/CkiYXoC78atKT1sI/hvqFA=
github.com/pt-main/pack v1.2.0 h1:OvUclASpNwkMw96Xr0mzU3x4FWVWCslKuiJOCKKkLPI=
github.com/pt-main/pack v1.2.0/go.mod h1:Se6SUhnOIQ4BblOVHVkOu9fUZxzdd7CSx36G/lNAhIA=
github.com/pt-main/tap v1.4.11 h1:BXMrfXN4ZfX9df6yjYrCYVev9JkIyxNKOns6IaJDczU=
github.com/pt-main/tap v1.4.11/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.4.13 h1:2KeJlw38nM4lOt6T3MZPuqTqp7UWmnGOeF1WfUbrXho=
github.com/pt-main/tap v1.4.13/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.4.14-0.20260922135441-c0b7d2c122cf h1:jGjEFT5E7psvoeHqQgaBx30Rua49UuPOv/m01A+EzAM=
github.com/pt-main/tap v1.4.14-0.20260922135441-c0b7d2c122cf/go.mod h1:M8UyfQ2yg3k/YZ9BUjoS8QYGdKu2l2aOEk3qwE2Qbgo=
github.com/pt-main/tap v1.4.14 h1:DbFwrdnu5yqOgIKw1RSDrLNXk68CFNXUlkpi8QqDmOs=
github.com/pt-main/tap v1.4.14/go.mod h1:ULQUJ/+8VIji9oq26pr1cmbXv+VUlhjsvq1n/vd4f3I=
github.com/pt-main/tap v1.5.3 h1:fHskNl95OLVg02F8c4LeF5W1FG0inPyKGMzZXSDon1s=
github.com/pt-main/tap v1.5.3/go.mod h1:N6lrtdlW90TqAp9pL08PjFchwcHQS62mjZybVUmHG6Y=
github.com/pt-main/tap/go v1.5.7 h1:GzPXBjJETJcXAeQkosjBzwrBCe3RLkKCUWumK/KdBWg=
github.com/pt-main/tap/go v1.5.7/go.mod h1:JhdaGAsrmcZVFSANOv8PlSuKcF0Uqx0Kad2XakiOPlM=
github.com/pt-main/tap/go v1.5.8 h1:f2thGjf2OgcIq3Ehtpt3Fln5ZLSm7q33+eodMP416/k=
github.com/pt-main/tap/go v1.5.8/go.mod h1:JhdaGAsrmcZVFSANOv8PlSuKcF0Uqx0Kad2XakiOPlM=
github.com/pt-main/tycl v1.3.8 h1:TU7/axhLYE7l3TGmbvth3AfGQOorwI5uDnSO8FxeD54=
github.com/pt-main/tycl v1.3.8/go.mod h1:zKmXw4/TpgHkz+gYo0D88HtofifVkcdcF3pOWSc4IIQ=
github.com/yuin/gopher-lua v1.1.2 h1:yF/FjE3hD65tBbt0VXLE13HWS9h34fdzJmrWRXwobGA=
github.com/yuin/gopher-lua v1.1.2/go.mod h1:7aRmXIWl37SqRf0koeyylBEzJ+aPt8A+mmkQ4f1ntR8=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

---

# main.go

```go
package run

var Version = "1.3.4"
```

---

# rubytempl.txt

```txt
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

---

# run/api/funcs.go

```go
package api

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-shellwords"
	localmode "github.com/pt-main/run/run/api/localMode"
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
		panic("Config dir path finding:" + err.Error())
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

func NewScriptConfig(name, script, description, ext string, tags []string) *shared.Config {
	conf := shared.NewNilConfig()
	conf.StringV["name"] = name
	conf.StringV["script"] = script
	conf.StringV["description"] = description
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

# run/api/localMode/main.go

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

# run/api/lua.go

```go
package api

import (
	"log"
	"path/filepath"
	"sync"
	"sync/atomic"

	lua "github.com/yuin/gopher-lua"
)

var GlobalFuncs = map[string]lua.LGFunction{}

func RegisterLuaFunc(name string, fun lua.LGFunction) {
	GlobalFuncs[name] = fun
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

	var (
		activeScripts int32
		wg            sync.WaitGroup
	)

	L.SetGlobal("run_script_parallel", L.NewFunction(func(L *lua.LState) int {
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

		atomic.AddInt32(&activeScripts, 1)
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer atomic.AddInt32(&activeScripts, -1)
			defer func() {
				if r := recover(); r != nil {
					log.Printf("script %q panicked: %v", name, r)
				}
			}()

			if err := RunScript(cfg, name, scriptArgs); err != nil {
				log.Printf("script %q failed: %v", name, err)
			}
		}()
		return 0
	}))

	L.SetGlobal("wait", L.NewFunction(func(L *lua.LState) int {
		wg.Wait()
		return 0
	}))

	for name, fun := range GlobalFuncs {
		L.SetGlobal(name, L.NewFunction(fun))
	}

	return L
}
```

---

# run/api/main.go

```go
package api

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	var errI core.ErrorInterface
	cfg, errI := tycl.Process(file, TyclContract, true)
	if errI != nil {
		return cfg, fmt.Errorf(format.FormatError(errI))
	}
	if _, ok := cfg.InnerArrV["templates"]; !ok {
		cfg.InnerArrV["templates"] = []*shared.Config{}
		err = UpdateConfig(cfg)
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func AddScript(conf *shared.Config, script, rawScriptName, scriptName, docs string, force bool) error {
	rawScriptName = scriptName + "_" + strings.ReplaceAll(rawScriptName, "/", "__")

	runScript := ""

	newScripts := []*shared.Config{}
	for _, script := range conf.InnerArrV["scripts"] {
		name := script.StringV["script"]
		if name == scriptName && !force {
			return fmt.Errorf("Can't add script: script already added. Use --force to replace script.")
		}
		if name != scriptName {
			newScripts = append(newScripts, script)
		}
	}
	conf.InnerArrV["scripts"] = newScripts
	addScript := true

	processed := false
	ext := filepath.Ext(rawScriptName)

	if strings.HasSuffix(rawScriptName, ".nd.task.lua") { // nd - no deps
		processed = true
		ext = ".nd.task.lua"
		runScript = TalRunScriptTemplate(rawScriptName, false)
	} else if strings.HasSuffix(rawScriptName, ".task.lua") {
		processed = true
		ext = ".task.lua"
		runScript = TalRunScriptTemplate(rawScriptName, true)
	}

	var fallback *string

	parseTempl := func(templ string) error {
		tpl, err := template.New(ext).Parse(templ)
		if err != nil {
			return fmt.Errorf("Add script: parsing extension template: %v", err)
		}
		var b strings.Builder
		err = tpl.Execute(&b, map[string]string{"ext": ext, "name": rawScriptName})
		if err != nil {
			return fmt.Errorf("Add script: executing extension template: %v", err)
		}
		runScript = b.String()
		return nil
	}

	if !processed {
		templs := conf.InnerArrV["templates"]
		for _, cfg := range templs {
			ext := cfg.StringV["ext"]
			templ := cfg.StringV["template"]

			if ext == "" {
				fallback = &templ
			}

			if strings.HasSuffix(rawScriptName, ext) && ext != "" {
				if err := parseTempl(templ); err != nil {
					return err
				}
				processed = true
			}
		}
	}

	if fallback != nil && !processed {
		if err := parseTempl(*fallback); err != nil {
			return err
		}
		processed = true
	}

	if !processed {
		switch ext {
		case ".py":
			processed = true
			runScript = PythonRunScriptTemplate(rawScriptName)
		case ".sh":
			processed = true
			runScript = BashRunScriptTemplate(rawScriptName)
		case ".bat":
			processed = true
			runScript = BatRunScriptTemplate(rawScriptName)
		case ".lua":
			processed = true
			runScript = script
			addScript = false
		}
	}
	if !processed {
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}
	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"], NewScriptConfig(scriptName, scriptName, docs, ext, nil))
	if err := NewRunScript(scriptName, runScript); err != nil {
		return err
	}
	if addScript {
		fpsplit := strings.Split(rawScriptName, "/")
		file := fpsplit[0]
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
	return nil
}

func AddTemplate(conf *shared.Config, ext, template string, force bool) error {
	templates := conf.InnerArrV["templates"]
	templatesA := []*shared.Config{}
	for _, templ := range templates {
		tExt := templ.StringV["ext"]
		if tExt == ext && !force {
			return fmt.Errorf("Can't add template: extension duplicate and has no force flag")
		} else {
			templatesA = append(templatesA, templ)
		}
	}
	c := shared.NewNilConfig()
	c.StringV["ext"] = ext
	res := strconv.Quote(template)
	c.StringV["template"] = res
	templatesA = append(templatesA, c)
	conf.InnerArrV["templates"] = templatesA
	return nil
}

func RemoveTemplate(conf *shared.Config, ext string) error {
	templates := conf.InnerArrV["templates"]
	templatesA := []*shared.Config{}
	for _, templ := range templates {
		tExt := templ.StringV["ext"]
		if tExt == ext {
			templatesA = append(templatesA, templ)
		}
	}
	conf.InnerArrV["templates"] = templatesA
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

func Upconf(conf *shared.Config, err error) error {
	if err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	return nil
}
```

---

# run/api/stdlib.go

```go
package api

import (
	"github.com/pt-main/tycl/shared"
)

func StdLib() (string, error) {
	config := shared.NewNilConfig()
	NewRunScript("test", 'print("test script"); print(script_path("test.py")); print(get_args()[1])')
	config.InnerArrV["scripts"] = append(config.InnerArrV["scripts"], NewScriptConfig(
		"test", "test", "[?BBK]Simple script for functions test[?RT]", "", []string{"__test"},
	))
	config.InnerArrV["templates"] = []*shared.Config{}
	conf, err := FormatConfig(config)
	if err != nil {
		return "", err
	}
	return conf, nil
}
```

---

# run/api/templates.go

```go
package api

import "fmt"

func TalRunScriptTemplate(name string, depsEnabled bool) string {
	return fmt.Sprintf('-- === CONFIGURATION ===
local task_name = script_path("%s")
local args = get_args()
local deps_enabled = %v
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "tal run "
if deps_enabled then
    cmd = cmd .. "--deps="" "
end
cmd = cmd .. escape(task_name)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = cli(cmd)
os.exit(result or 0)', name, depsEnabled)
}

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

# run/api/tycl.go

```go
package api

var TyclContract = '
flexible {
	scripts: objects = strict {
		name: string,
		script: string,
		description: string,
		tags: strings,
		ext: string,
	},
	templates: objects = strict {
		ext: string,
		template: string,
	},
}
'
```

---

# run/runcli/cli.go

```go
package runcli

import (
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run/run/api"
	runlib "github.com/pt-main/run/run/runcli/handlers"
	luaruntime "github.com/pt-main/run/tal/lua"
	"github.com/pt-main/run/tal/runtime"
	tap "github.com/pt-main/tap/go"
	lua "github.com/yuin/gopher-lua"
)

func NewCli() (*tap.Parser, error) {
	var lp *tap.Parser
	conf := tap.DefaultParserConfig()
	conf.BuiltinVerboseDebug = true
	p := tap.NewParser("run", '[?BE]╭─────── [?BRD]Run[?RT]
[?BE]⎬─ [?RT]Simple and powerful script manager
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]', []string{"-h", "-help"}, conf)

	p.AddSubcommand("tal", runtime.CreateCli())

	p.AddCommand("-r",
		runlib.MakeRunHandler(true),
		'[?GN]Run a script by name.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -r <name> [args...] [--args="..."][[?RT]
  [?BBK]run -r --tagged="..." [--parallel] [--args="..."][?RT]
[?YW]Flags:[?RT]
  [?GN]--tagged="tag1;tag2;..."[?RT]  Run all scripts with any of the specified tags (semicolon-separated)
  [?GN]--parallel[?RT]                Run tagged scripts in parallel
  [?GN]--args="..."[?RT]              Explicitly pass arguments to script (useful when args conflict with run flags)
  [?GN]--args[?RT]                    Pass no arguments (instead of passing run flags)
[?YW]Examples:[?RT]
  [?BBK]run -r deploy --env=prod[?RT]
  [?BBK]run -r --tagged="deploy;test" --parallel[?RT]
  [?BBK]run -r --tagged="build" --args="--verbose"[?RT]',
		nil, nil, true)

	p.AddCommand(tap.DEFAULT_CMD,
		runlib.MakeRunHandler(false),
		'[?GN]Run a script by name (when name doesn't conflict with run commands).[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run <script> [args...][?RT]
  [?BBK]run <cmd> [args...][?RT]
[?YW]Examples:[?RT]
  [?BBK]run deploy --env=prod[?RT]
  [?BBK]run mypy arg1 arg2[?RT]',
		nil, nil, true)

	m, err := NewManage()
	if err != nil {
		return nil, err
	}
	p.AddSubcommand("manage", m)

	p.AddSubcommand("sys", NewSys())

	runcli := func(L *lua.LState) int {
		input := L.OptString(1, "")
		if input == "" {
			L.Push(lua.LString("missing command string"))
			return 2
		}
		parsed, err := shellwords.Parse(input)
		if err != nil {
			fmt.Println("Parsing cli args:", err)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		if lp == nil {
			lp, err = NewCli()
			if err != nil {
				L.Push(lua.LString(err.Error()))
				return 2
			}
		}
		if err := Process(lp, parsed); err != nil {
			fmt.Println("Run cli err:", err)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		return 0
	}

	luaruntime.RegisterLuaFunc("run_cli", func(changedFiles, args []string) lua.LGFunction {
		return runcli
	})

	api.RegisterLuaFunc("run_cli", runcli)

	return p, nil
}
```

---

# run/runcli/handlers/handlers.go

```go
package runlib

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/pt-main/run/run/api"
	tap "github.com/pt-main/tap/go"
	"github.com/pt-main/tap/go/color"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func AddHandler(p *tap.Parser, s []string) error {
	_, force := p.Flags["force"]

	script, err := utils.OpenF(s[0])
	if err != nil {
		return err
	}
	docs := ""
	if len(s) > 2 {
		docs = s[2]
	}
	conf, err := api.GetCfg()
	if err != nil {
		return err
	}
	return api.Upconf(conf, api.AddScript(conf, script, s[0], s[1], docs, force))
}

func RemoveHandler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
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
	return api.UpdateConfig(cfg)
}

func ListHandler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
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

func MakeRunHandler(hasRawArgs bool) func(p *tap.Parser, s []string) error {
	return func(p *tap.Parser, s []string) error {
		cfg, err := api.GetCfg()
		if err != nil {
			return err
		}
		idx := 0
		if len(s) > 0 {
			idx += 1
		}
		if len(p.RawArgs) > 0 && slices.Contains([]string{"--gm", "--globalmode",
			"--lm", "--localmode"}, p.RawArgs[0]) {
			idx += 1
		}
		if hasRawArgs {
			idx += 1
		}
		var args []string = nil
		args_, ok := p.Flags["args"]
		if ok {
			args, err = api.ProcessShell(args_)
			if err != nil {
				return err
			}
			if args == nil {
				args = []string{}
			}
		}
		if args == nil && len(p.RawArgs) > 0 {
			args = p.RawArgs[idx:]
		}
		if tags_, ok := p.Flags["tagged"]; ok {
			_, parallel := p.Flags["parallel"]
			tags := strings.Split(tags_, ";")
			errs := []string{}
			var errsMu sync.Mutex
			var wg sync.WaitGroup

			for _, script := range cfg.InnerArrV["scripts"] {
				scrTags := script.StringArrV["tags"]
				scriptName := script.StringV["name"]
				for _, tag := range scrTags {
					if slices.Contains(tags, tag) {
						p.Print("verbose", "Run %v: ", scriptName)
						if parallel {
							wg.Add(1)
							go func(name string) {
								defer wg.Done()
								if err := api.RunScript(cfg, name, args); err != nil {
									p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
									errsMu.Lock()
									errs = append(errs, err.Error())
									errsMu.Unlock()
								} else {
									p.Print("verbose", "[?GN]Ok[?RT]")
								}
							}(scriptName)
						} else {
							if err := api.RunScript(cfg, scriptName, args); err != nil {
								p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
								errs = append(errs, err.Error())
							} else {
								p.Print("verbose", "[?GN]Ok[?RT]")
							}
						}
					}
				}
			}

			wg.Wait()
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
			return api.RunScript(cfg, name, args)
		}
	}
}

func TagHahdler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	addTags := []string{}
	rmTags := []string{}
	for _, tag := range s[1:] {
		if strings.HasPrefix(tag, "!") {
			rmTags = append(rmTags, tag[1:])
		} else {
			addTags = append(addTags, tag)
		}
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			tags := append(script.StringArrV["tags"], addTags...)
			newTags := []string{}
			for _, tag := range tags {
				if !slices.Contains(rmTags, tag) {
					newTags = append(newTags, tag)
				}
			}
			script.StringArrV["tags"] = newTags
			break
		}
	}
	return api.UpdateConfig(cfg)
}
```

---

# run/runcli/handlers/install.go

```go
package runlib

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run/run/api"
	"github.com/pt-main/run/tal"
	tap "github.com/pt-main/tap/go"
)

func downloadScript(url string) (content string, fileName string, err error) {
	if strings.Contains(url, "github.com/") {
		rawURL, fname, err := parseGitHubURL(url)
		if err != nil {
			return "", "", err
		}
		url = rawURL
		fileName = fname
	} else {
		fileName = filepath.Base(url)
		if idx := strings.Index(fileName, "?"); idx != -1 {
			fileName = fileName[:idx]
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	return string(data), fileName, nil
}

func parseGitHubURL(rawURL string) (rawContentURL string, fileName string, err error) {

	u := rawURL
	if strings.HasPrefix(u, "https://") {
		u = strings.TrimPrefix(u, "https://")
	} else if strings.HasPrefix(u, "http://") {
		u = strings.TrimPrefix(u, "http://")
	}

	if !strings.HasPrefix(u, "github.com/") {
		return "", "", fmt.Errorf("not a GitHub URL")
	}
	u = strings.TrimPrefix(u, "github.com/")

	parts := strings.SplitN(u, "@", 2)
	if len(parts) == 2 {
		repo := parts[0]
		rest := parts[1]
		slashIdx := strings.Index(rest, "/")
		if slashIdx == -1 {
			return "", "", fmt.Errorf("missing path after ref")
		}
		ref := rest[:slashIdx]
		path := rest[slashIdx+1:]
		rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, path)
		fileName = filepath.Base(path)
		return rawContentURL, fileName, nil
	}

	idx := strings.Index(u, "/blob/")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid GitHub URL: missing '@' or '/blob/'")
	}
	repoPart := u[:idx]
	rest := u[idx+len("/blob/"):]
	slashIdx := strings.Index(rest, "/")
	if slashIdx == -1 {
		return "", "", fmt.Errorf("invalid blob URL: missing branch/path")
	}
	branch := rest[:slashIdx]
	path := rest[slashIdx+1:]
	rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repoPart, branch, path)
	fileName = filepath.Base(path)
	return rawContentURL, fileName, nil
}

func InstallHandler(p *tap.Parser, s []string) error {
	if len(s) < 1 {
		return fmt.Errorf("need at least URL")
	}

	url := s[0]
	scriptName := ""
	docs := ""

	if len(s) > 1 {
		scriptName = s[1]
	}
	if len(s) > 2 {
		docs = s[2]
	}

	content, rawName, err := downloadScript(url)
	if err != nil {
		return err
	}

	if scriptName == "" {
		ext := filepath.Ext(rawName)
		scriptName = strings.TrimSuffix(rawName, ext)
	}

	_, force := p.Flags["force"]

	// running tal isntallation file
	if rawName == "run.task.lua" {
		args := s[1:]
		_args, hasArgs := p.Flags["args"]
		if hasArgs {
			args, err = shellwords.Parse(_args)
			if err != nil {
				return fmt.Errorf("Parsing args: %v", err)
			}
		}

		err := tal.Process([]string{}, args, content)
		if err != nil {
			return err
		}
		return nil
	}

	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}

	// adding script
	return api.Upconf(cfg, api.AddScript(cfg, content, rawName, scriptName, docs, force))
}
```

---

# run/runcli/handlers/templates.go

```go
package runlib

import (
	"fmt"

	"github.com/pt-main/run/run/api"
	tap "github.com/pt-main/tap/go"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func TemplateAdd(p *tap.Parser, s []string) (err error) {
	template := ""
	if source, has := p.Flags["source"]; has {
		template = source
	} else if len(s) > 1 {
		template, err = utils.OpenF(s[1])
		if err != nil {
			return
		}
	} else {
		return fmt.Errorf("Can't add template: source is not provided")
	}

	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	_, force := p.Flags["force"]
	err = api.AddTemplate(cfg, s[0], template, force)
	// fmt.Println(force, template, p.Flags, p.RawArgs, s)
	if err != nil {
		return
	}
	return api.UpdateConfig(cfg)
}

func TemplateRem(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	templates := []*shared.Config{}
	for _, templ := range cfg.InnerArrV["templates"] {
		if templ.StringV["ext"] == s[0] {
			templates = append(templates, templ)
		}
	}
	cfg.InnerArrV["templates"] = templates
	return api.UpdateConfig(cfg)
}
```

---

# run/runcli/manage.go

```go
package runcli

import (
	runlib "github.com/pt-main/run/run/runcli/handlers"
	tap "github.com/pt-main/tap/go"
)

func NewManage() (p *tap.Parser, err error) {
	manageP := tap.NewParser("manage", "[?GN]Manage run data.[?RT]", []string{"help", "-help", "-h"}, tap.DefaultParserConfig())

	manageP.AddCommand("script-add", runlib.AddHandler,
		'[?GN]Add a local script to the configuration.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage script-add <path> <name> [description] [--force][?RT]
[?BBK]Supported extensions:[?RT] .py, .sh, .bat, .lua, .task.lua
[?YW]Flags:[?RT]
  [?GN]--force[?RT]    Replace existing script with the same name
[?YW]Examples:[?RT]
  [?BBK]run manage script-add ./deploy.py deploy "Deploy script"[?RT]
  [?BBK]run manage scradd ./build.sh build --force[?RT]',
		[]string{"path", "name"}, []string{"description"}, false)
	if err = manageP.AddAlias("scradd", "script-add"); err != nil {
		return
	}

	manageP.AddCommand("script-remove", runlib.RemoveHandler,
		'[?GN]Remove a script from the configuration.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage script-remove <name>[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage script-remove myscript[?RT]',
		[]string{"name"}, nil, false)
	if err = manageP.AddAlias("screm", "script-remove"); err != nil {
		return
	}

	manageP.AddCommand("templ-add", runlib.TemplateAdd,
		'[?GN]Add a custom template for a file extension.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage templ-add <ext> --source="..." [--force][?RT]
  [?BBK]run manage templ-add <ext> <file> [--force][?RT]
[?YW]Flags:[?RT]
  [?GN]--source="..."[?RT]  Template content as string (instead of file)
  [?GN]--force[?RT]         Overwrite existing template for this extension
[?YW]Examples:[?RT]
  [?BBK]run manage templ-add ".go" templ.txt --force[?RT]
  [?BBK]run manage templ-add ".go" --source="..."[?RT]',
		[]string{"ext"}, []string{"file"}, false)
	if err = manageP.AddAlias("tladd", "templ-add"); err != nil {
		return
	}

	manageP.AddCommand("templ-remove", runlib.TemplateRem,
		'[?GN]Remove a template for an extension.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage templ-remove <ext>[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage templ-rem ".go"[?RT]',
		[]string{"ext"}, nil, false)
	if err = manageP.AddAlias("tlrem", "templ-remove"); err != nil {
		return
	}

	manageP.AddCommand("tag",
		runlib.TagHahdler,
		'[?GN]Add or remove tags for a script.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage tag <script> <tag1> <tag2> ... [?RT]
  Prefix tag with '!' to remove it.
[?YW]Examples:[?RT]
  [?BBK]run manage tag deploy prod staging[?RT]
  [?BBK]run manage tag deploy !prod[?RT]',
		[]string{"script"}, []string{"tag"}, true)

	manageP.AddCommand("list", runlib.ListHandler,
		'[?GN]List all registered scripts.[?RT]
[?BBK]Shows script names, descriptions, and tags.[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage list[?RT]',
		nil, nil, false)

	manageP.AddCommand("install", runlib.InstallHandler,
		'[?GN]Download and install a script from any URL.[?RT]
[?BBK]Supported URLs:[?RT]
  - Raw file URLs [?BBK](https://raw.githubusercontent.com/...)[?RT]
  - GitHub blob URLs [?BBK](github.com/user/repo/blob/branch/path/script.py)[?RT]
  - GitHub simpler URLs [?BBK](github.com/user/repo@branch/path/script.py)[?RT]
  - Running installation script [?BBK](github.com/user/repo@branch/path/run.task.lua)[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage install <url> [name] [description] [--force] [--args="..."][?RT]
[?YW]Flags:[?RT]
  [?GN]--force[?RT]        Replace existing script with the same name
  [?GN]--args="..."[?RT]   Pass args to run script (for run.task.lua)
[?YW]Examples:[?RT]
  [?BBK]run manage install https://raw.githubusercontent.com/user/repo/main/deploy.py[?RT]
  [?BBK]run manage install github.com/user/repo@branch/script.py myscript[?RT]
  [?BBK]run manage install github.com/user/repo@branch/run.task.lua[?RT]',
		[]string{"url"}, []string{"name", "description"}, false)

	return manageP, nil
}
```

---

# run/runcli/process.go

```go
package runcli

import (
	"fmt"

	"github.com/pt-main/run/run/api"
	localmode "github.com/pt-main/run/run/api/localMode"
	tap "github.com/pt-main/tap/go"
)

func Process(cli *tap.Parser, args []string) error {
	lm := localmode.IsLocalmode()
	temp := lm

	if len(args) > 0 {
		if args[0] == "--localmode" || args[0] == "--lm" {
			temp = true
		} else if args[0] == "--globalmode" || args[0] == "--gm" {
			temp = false
		}
	}
	localmode.Set(temp)

	defer func() {
		if localmode.IsLocalmode() == temp {
			localmode.Set(lm)
		}
	}()

	ok, err := api.CheckConfigDir()
	if err != nil {
		return fmt.Errorf("Can't check installation: %v", err)
	}
	if !ok {
		if err := api.InstallConfigDir(); err != nil {
			return fmt.Errorf("Can't make run dir: %v", err)
		}
	}

	err = cli.Parse(args)
	fmt.Println(cli.RawArgs, cli.Flags)
	if err != nil {
		return err
	}
	return nil
}
```

---

# run/runcli/sys.go

```go
package runcli

import (
	"fmt"

	"github.com/pt-main/run"
	"github.com/pt-main/run/run/api"
	localmode "github.com/pt-main/run/run/api/localMode"
	"github.com/pt-main/run/tal"
	tap "github.com/pt-main/tap/go"
)

func NewSys() *tap.Parser {
	sysP := tap.NewParser("sys", "[?GN]Run systems.[?RT]", []string{"help", "-help", "-h"}, tap.DefaultParserConfig())

	sysP.AddCommand("version", func(p *tap.Parser, s []string) error {
		fmt.Println("run v" + run.Version)
		fmt.Println("tal v" + tal.Version)
		fmt.Println("humanmade, by Pt, Apache 2.0 licence")
		return nil
	},
		'[?GN]Show version and license information.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run sys version[?RT]',
		nil, nil, false)
	if err := sysP.AddAlias("-v", "version"); err != nil {
		panic("SYSTEM ERROR: CREATING CLI: " + err.Error())
	}

	sysP.AddCommand("localmode", func(p *tap.Parser, s []string) error {
		if len(s) == 0 {
			fmt.Println("localmode:", localmode.IsLocalmode(), "| path:", api.ConfigDirPath())
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
	},
		'[?GN]Set or show the current working mode (global/local).[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run sys localmode[?RT]           Show current mode and config path
  [?BBK]run sys localmode true[?RT]      Enable local mode (use .run/ in current directory)
  [?BBK]run sys localmode false[?RT]     Disable local mode (use ~/run/)
[?YW]Examples:[?RT]
  [?BBK]run sys localmode[?RT]
  [?BBK]run sys localmode true[?RT]',
		nil, []string{"mode"}, false)
	if err := sysP.AddAlias("-lm", "localmode"); err != nil {
		panic("SYSTEM ERROR: CREATING CLI: " + err.Error())
	}

	return sysP
}
```

---

# tal/README-ru.md

```md
# tal - инкрементальный таскер с Lua и зависимостями

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal - Task Lua

'''bash
go install github.com/pt-main/run/cmd/tal@latest
'''

**tal** - это простой современный таскер с Lua в роли языка скриптования. Он позволяет описывать задачи на обычном Lua с аннотациями, отслеживать изменения файлов и запускать только то, что действительно изменилось.

---

## Зачем tal?

| Проблема | tal решает |
|----------|------------|
| **Makefile сложно читать и писать** | Простой DSL с комментариями и Lua вместо Shell |
| **Инкрементальность работает криво** | Хеши SHA256 вместо времени модификации |
| **Нет вызова задач друг из друга** | Можно вызывать таски через встроенную функцию |
| **Зависимости от файлов громоздкие** | '-- #depends file1 file2' работает из коробки |

tal даёт **инкрементальность, простоту и Lua** - всё в одном инструменте.

---

## Установка

'''bash
go install github.com/pt-main/run/cmd/tal@latest
'''

Также при установке [run](https://github.com/pt-main/run) tal устанавливается автоматически и доступен как 'run tal ...'.

При первом запуске 'tal update' создаст '.tal.pack' - файл с хешами файлов в текущей директории. Инициализировать проект (создать '.tal.pack' и 'main.task.lua') можно командой 'tal init'.

---

## Синтаксис

Файл тасков пишется на обычном Lua с аннотациями в комментариях, не нарушая синтаксис.

### Основные конструкции

| Конструкция | Описание |
|-------------|----------|
| '-- @taskname' | Начало блока задачи |
| '-- @' | Главный блок (запускается по умолчанию) |
| '-- @!' | Глобальный блок (выполняется до main и регистрации тасков) |
| '-- #depends <glob...>' | Команда-зависимость от файлов (проверяются по хешам). Пути задаются в формате glob |

Любой другой код считается обычным Lua-кодом. Файл обязан начинаться с глобального или обычного блока.

Пример:

'''lua
-- @!
-- Глобальный код: переменные, функции, импорты
local function log(msg) print("[TASK] " .. msg) end

-- @build
-- #depends *.go
log("Building...")
os.execute("go build .")

-- @test
-- #depends test/**
log("Testing...")
os.execute("go test .")

-- @
-- Запускается по умолчанию
script("build")
script("test")

update() -- обновляет .tal.pack, подтверждая, что все изменения обработаны
'''

---

## Команды CLI

| Команда | Описание | Пример |
|---------|----------|--------|
| 'tal run <file> [args...] [--deps="..."]' | Разбирает '<file>' и выполняет DSL с аргументами, используя '.tal.pack' (обязателен, если не передан '--deps') | 'tal run main.task.lua build' |
| 'tal update' | Обновить или принудительно инициализировать '.tal.pack' | 'tal update' |
| 'tal init' | Инициализировать проект (создаёт '.tal.pack' и 'main.task.lua') | 'tal init' |
| 'tal list <file> [file...]' | Показать все задачи в файле (включая Global и Main) | 'tal list main.task.lua' |

Чтобы получить больше информации:

'''bash
tal help
'''

### Флаг '--deps'

'tal run ...' поддерживает флаг '--deps="dep1;dep2"', который позволяет передать в скрипт зависимости для аннотации '-- #depends' (разделитель - ';'). При отсутствии флага скрипт работает с '.tal.pack'.

'''bash
tal run main.task.lua build --deps="main.go;go.mod"
'''

Аргументы после имени файла передаются напрямую в Lua через 'get_args()' и не интерпретируются CLI.

---

## Как работает инкрементальность

Инкрементальность включается командой 'depends' ('-- #depends ...') и без неё не работает.

1. 'tal' сканирует текущую директорию и вычисляет SHA256 для всех файлов.
2. Хеши сохраняются в '.tal.pack' (бинарный формат, использует ['pack'](https://github.com/pt-main/pack)).
3. При следующем запуске 'tal' сравнивает хеши, определяет, какие файлы изменились, и автоматически обновляет хеши.
4. В сгенерированном Lua-скрипте массив 'changed_list' содержит пути к изменённым файлам.
5. Рантайм проверяет зависимости каждой задачи и выполняет только те, у которых изменился хотя бы один зависимый файл.

Для обновления '.tal.pack' обязательно использовать функцию 'update()'.

---

## Встроенный рантайм Lua

Каждая задача - это Lua-функция, которая выполняется в окружении с доступом к функциям:

- 'changed_list' - таблица с путями изменённых файлов (относительно текущей директории).
- 'get_args()' - таблица аргументов, переданных в 'tal run'.
- 'script(name)' - выполнение скрипта.
- 'shell(string)' - сокращение 'os.execute'.
- 'print_colored(string)' - цветной вывод (использует систему цветов из ['tap'](https://github.com/pt-main/tap).color).
- 'update()' - пересчитывает хеши и полностью обновляет '.tal.pack'.
- 'match_pattern(glob, path)' - проверяет, соответствует ли путь 'path' шаблону 'glob' (поддерживается синтаксис doublestar).

Когда tal используется из 'run cli', становится доступна дополнительная функция - 'run_cli(args_string)', которая напрямую вызывает run и парсит аргументы из строки на входе.

**Важно**: нельзя использовать внешние Lua-библиотеки (интерпретатор Lua в tal написан на [Go](https://github.com/yuin/gopher-lua) и не зависит от системы и установленных Lua-библиотек).

---

## Структура проекта

'''
.
├── main.task.lua      # файл с задачами (DSL)
├── .tal.pack          # бинарный файл с хешами (создаётся автоматически при запуске tal)
└── ...
'''

---

## Сравнение с аналогами

| Возможность | tal | make | just | task |
|-------------|-----|------|------|------|
| **Инкрементальность по хешам** | Да | Нет | Нет | Да |
| **Язык скриптов** | Lua с аннотациями | Shell | Shell | Shell |
| **Вызов других задач** | Да | Да | Нет | Да |
| **Простота написания** | Просто | Сложно | Просто | Средне |

---

## Лицензия

Apache 2.0 - подробности в [LICENSE](LICENSE).

---

By Pt, 2026 - написано с использованием 'lc', 'tap', 'pack'.
```

---

# tal/README.md

```md
# tal - incremental task runner with Lua and dependencies

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal - Task Lua

'''bash
go install github.com/pt-main/run/cmd/tal@latest
'''

**tal** is a simple modern task runner with Lua as its scripting language. It lets you describe tasks in plain Lua with annotations, track file changes, and run only what has actually changed.

---

## Why tal?

| Problem | tal solves |
|----------|------------|
| **Makefiles are hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incrementality works poorly** | SHA256 hashes instead of modification time |
| **No calling tasks from one another** | Tasks can be called via a built-in function |
| **File dependencies are cumbersome** | '-- #depends file1 file2' works out of the box |

tal gives you **incrementality, simplicity, and Lua** - all in one tool.

---

## Installation

'''bash
go install github.com/pt-main/run/cmd/tal@latest
'''

Also, when [run](https://github.com/pt-main/run) is installed, tal is installed automatically and is available as 'run tal ...'.

On first run, 'tal update' will create '.tal.pack' - a file with hashes of files in the current directory. You can initialize a project (create '.tal.pack' and 'main.task.lua') with the 'tal init' command.

---

## Syntax

The task file is written in plain Lua with annotations in comments, without breaking the syntax.

### Basic constructs

| Construct | Description |
|-------------|-------------|
| '-- @taskname' | Start of a task block |
| '-- @' | Main block (runs by default) |
| '-- @!' | Global block (executed before main and task registration) |
| '-- #depends <glob...>' | File dependency command (checked by hashes). Paths are specified in glob format |

Any other code is considered regular Lua code. The file must start with a global or regular block.

Example:

'''lua
-- @!
-- Global code: variables, functions, imports
local function log(msg) print("[TASK] " .. msg) end

-- @build
-- #depends *.go
log("Building...")
os.execute("go build .")

-- @test
-- #depends test/**
log("Testing...")
os.execute("go test .")

-- @
-- Runs by default
script("build")
script("test")

update() -- updates .tal.pack, confirming that all changes have been processed
'''

---

## CLI commands

| Command | Description | Example |
|---------|-------------|---------|
| 'tal run <file> [args...] [--deps="..."]' | Parses '<file>' and executes the DSL with arguments, using '.tal.pack' (required if '--deps' is not passed) | 'tal run main.task.lua build' |
| 'tal update' | Update or force-initialize '.tal.pack' | 'tal update' |
| 'tal init' | Initialize a project (creates '.tal.pack' and 'main.task.lua') | 'tal init' |
| 'tal list <file> [file...]' | Show all tasks in the file (including Global and Main) | 'tal list main.task.lua' |

To get more information:

'''bash
tal help
'''

### The '--deps' flag

'tal run ...' supports the '--deps="dep1;dep2"' flag, which allows passing dependencies to the script for the '-- #depends' annotation (separator is ';'). If the flag is absent, the script works with '.tal.pack'.

'''bash
tal run main.task.lua build --deps="main.go;go.mod"
'''

Arguments after the file name are passed directly to Lua via 'get_args()' and are not interpreted by the CLI.

---

## How incrementality works

Incrementality is enabled by the 'depends' command ('-- #depends ...') and does not work without it.

1. 'tal' scans the current directory and computes SHA256 for all files.
2. Hashes are saved in '.tal.pack' (binary format, uses ['pack'](https://github.com/pt-main/pack)).
3. On the next run, 'tal' compares hashes, determines which files changed, and automatically updates the hashes.
4. In the generated Lua script, the 'changed_list' array contains paths to changed files.
5. The runtime checks each task's dependencies and executes only those for which at least one dependent file has changed.

To update '.tal.pack', you must use the 'update()' function.

---

## Built-in Lua runtime

Each task is a Lua function that runs in an environment with access to the functions:

- 'changed_list' - table with paths of changed files (relative to the current directory).
- 'get_args()' - table of arguments passed to 'tal run'.
- 'script(name)' - executing a script.
- 'shell(string)' - shorthand for 'os.execute'.
- 'print_colored(string)' - colored output (uses the color system from ['tap'](https://github.com/pt-main/tap).color).
- 'update()' - recalculates hashes and completely updates '.tal.pack'.
- 'match_pattern(glob, path)' - checks whether path 'path' matches the 'glob' pattern (doublestar syntax is supported).

When tal is used from 'run cli', an additional function becomes available - 'run_cli(args_string)', which directly calls run and parses arguments from the input string.

**Important**: you cannot use external Lua libraries (the Lua interpreter in tal is written in [Go](https://github.com/yuin/gopher-lua) and does not depend on the system or installed Lua libraries).

---

## Project structure

'''
.
├── main.task.lua      # task file (DSL)
├── .tal.pack          # binary file with hashes (created automatically when tal runs)
└── ...
'''

---

## Comparison with alternatives

| Feature | tal | make | just | task |
|-------------|-----|------|------|------|
| **Hash-based incrementality** | Yes | No | No | Yes |
| **Scripting language** | Lua with annotations | Shell | Shell | Shell |
| **Calling other tasks** | Yes | Yes | No | Yes |
| **Ease of writing** | Easy | Hard | Easy | Medium |

---

## License

Apache 2.0 - details in [LICENSE](LICENSE).

---

By Pt, 2026 - written using 'lc', 'tap', 'pack'.
```

---

# tal/core/main.go

```go
package core

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/iancoleman/orderedmap"
	"github.com/pt-main/pack/lib/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
)

// FileHash computes SHA256 hash of a file in a streaming fashion.
func FileHash(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// SaveState walks through the directory 'where' recursively,
// computes SHA256 hash for each file, and returns an ordered map
// where key = absolute file path, value = hex-encoded hash.
//
// Uses parallel workers and streaming reads for better performance.
func SaveState(where string) (*orderedmap.OrderedMap, error) {
	const maxWorkers = 32 // can be adjusted or set to runtime.NumCPU()

	abs, err := filepath.Abs(where)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return nil, os.ErrInvalid
	}

	var files []string
	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := os.Stat(path)
		if err != nil {
			return nil
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	type result struct {
		path string
		hash []byte
		err  error
	}
	results := make(chan result, len(files))
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	for _, p := range files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			h, err := FileHash(path)
			results <- result{path, h, err}
		}(p)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var entries []result
	for r := range results {
		if r.err != nil {
			return nil, r.err
		}
		entries = append(entries, r)
	}

	// Sort for deterministic order in orderedmap
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].path < entries[j].path
	})

	om := orderedmap.New()
	for _, e := range entries {
		om.Set(e.path, e.hash)
	}
	return om, nil
}

// Changes compares a previous state (was) with the current state of the directory 'where'.
// It returns a list of absolute file paths that are either new or modified.
// The comparison is based on SHA256 hashes.
func Changes(was *orderedmap.OrderedMap, where string) ([]string, error) {
	now, err := SaveState(where)
	if err != nil {
		return nil, err
	}

	// Build a map for O(1) lookup of old hashes
	wasMap := make(map[string][]byte, len(was.Keys()))
	for _, k := range was.Keys() {
		v, _ := was.Get(k)
		wasMap[k] = v.([]byte)
	}

	var res []string
	for _, k := range now.Keys() {
		oldHash, exists := wasMap[k]
		if !exists {
			res = append(res, k)
			continue
		}
		newHash, _ := now.Get(k)
		if !bytes.Equal(oldHash, newHash.([]byte)) {
			res = append(res, k)
		}
	}
	return res, nil
}

func StateAsPackCore(data *orderedmap.OrderedMap) ([]byte, error) {
	c := core.NewCore(data)
	res, err := c.CreateFile()
	if err != nil {
		err = errors.New(lang.GetRealErrorReverse(err))
	}
	return res, err
}

func PackCoreAsState(data []byte) (*orderedmap.OrderedMap, error) {
	c := core.NewCore(nil)
	err := c.ReadFile(data)
	if err != nil {
		err = errors.New(lang.GetRealErrorReverse(err))
	}
	return c.Containers, err
}

func Update() error {
	st, err := SaveState(".")
	if err != nil {
		return err
	}
	file, err := StateAsPackCore(st)
	if err != nil {
		return err
	}
	return Write(shared.TalFile, file)
}

func OpenF(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Open: %v", err)
	}
	return string(data), nil
}

func Open(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Open: %v", err)
	}
	return data, nil
}

func Write(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	return nil
}
```

---

# tal/generation/generate.go

```go
package generation

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
)

func GenerateCode(from *lang.TalCode) (string, core.ErrorInterface) {
	t := &Tasker{
		Functions: make([]string, 0),
	}
	res := ""
	res += "-- ==== RUNTIME CODE ==== --\n"
	res += GetRuntime() + "\n"
	if from.Global != nil {
		res += "\n-- ==== GLOBAL CODE ==== --\n"
		res += from.Global.Code + "\n"
	}
	for _, task := range from.Blocks {
		err := t.GenerateTask(task)
		if err != nil {
			return "", err
		}
	}
	res += "\n-- ==== TASKS DECLARATION ==== --\n"
	res += strings.Join(t.Functions, "\n\n")
	if from.Main != nil {
		res += "\n\n-- ==== MAIN CODE ==== --\n"
		res += from.Main.Code + "\n"
	}
	return res, nil
}

type Tasker struct {
	Functions []string
}

func (t *Tasker) GenerateTask(ts *lang.TalSection) (err core.ErrorInterface) {
	res := ""
	var patterns []string
	for cmd, args := range ts.Cmds {
		switch cmd {
		case "depends":
			parts := strings.Fields(args)
			patterns = append(patterns, parts...)
			continue
		default:
			err = core.Err(shared.GenerationError, "Unknown cmd")
		}
		return core.Wrap(shared.GenerationError, err, "Error in '%v' cmd", cmd)
	}

	patternsLua := "{"
	for i, p := range patterns {
		if i > 0 {
			patternsLua += ", "
		}
		patternsLua += fmt.Sprintf("%q", p)
	}
	patternsLua += "}"

	res += fmt.Sprintf('tasker.add(%v, 
"%v", function()
%v
end)', patternsLua, ts.Name, ts.Code)
	t.Functions = append(t.Functions, res)
	return nil
}
```

---

# tal/generation/runtime.go

```go
package generation

func GetRuntime() string {
	return '---@type string[]
local changed_list = changed()

---@class Task
---@field deps string[]
---@field func fun()

---@class Tasker
---@field tasks table<string, Task>

local tasker = {
    tasks = {}
}

---@param deps string[]
---@param changed string[]
---@return boolean
function tasker.has_any_dep_changed(patterns, changed)
    for _, pat in ipairs(patterns) do
        for _, ch in ipairs(changed) do
            if match_pattern(pat, ch) then
                return true
            end
        end
    end
    return false
end

---@param deps string[]
---@param name string
---@param func fun()
function tasker.add(deps, name, func)
    tasker.tasks[name] = {
        deps = deps,
        func = func
    }
end

---@param name string
function tasker.run(name)
    if tasker.tasks[name] == nil then
        error("Has not task: '" .. name .. "'")
        return
    end
    local task = tasker.tasks[name]
    if #task.deps == 0 then
        task.func()
        return
    end
    if tasker.has_any_dep_changed(task.deps, changed_list) then
        task.func()
    end
end

-- for simple call in scripts
function script(name) 
    tasker.run(name)
end

function shell(input)
    os.execute(input)
end'
}
```

---

# tal/lang/errFmt.go

```go
package lang

import (
	goerr "errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/run/tal/shared"
	"github.com/pt-main/tap/go/color"
)

var (
	// whereSpace is a colored prefix for indentation (blue line).
	whereSpace = color.Set("  [?BE]|[?RT]")
	// whereRedSpace is a colored prefix for error context (red line).
	whereRedSpace = color.Set("  [?RD]|[?RT]")
	// errSpace is a colored prefix for error messages (red arrow).
	errSpace = color.Set(" [?RD]>[?RT] ")
)

func GetRealErrorReverse(err error) string {
	if err == nil {
		return ""
	}

	if ce, ok := err.(core.ErrorInterface); ok {
		return ce.Format()
	}

	var parts []string
	cur := err
	for cur != nil {

		if ce, ok := cur.(core.ErrorInterface); ok {

			innerFormatted := ce.Format()
			if len(parts) > 0 {

				outerMsg := strings.Join(reverse(parts), ": ")
				return outerMsg + ": " + innerFormatted
			}
			return innerFormatted
		}

		parts = append(parts, cur.Error())

		next := goerr.Unwrap(cur)
		if next == nil {
			break
		}
		cur = next
	}

	if len(parts) > 0 {
		return strings.Join(reverse(parts), ": ")
	}
	return err.Error()
}

func reverse(s []string) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}

// GetErr extracts the inner error if it is a core.ErrorInterface.
// If the inner error is not a core.ErrorInterface, it wraps it into a core.Error
// with SystemError code.
func GetErr(ei core.ErrorInterface) core.ErrorInterface {
	inner := ei.Unwrap()
	if inner == nil {
		return nil
	}
	res, ok := inner.(core.ErrorInterface)
	if !ok {
		res = &core.Error{
			Code:  shared.SystemError,
			Msg:   inner.Error(),
			Meta:  make(map[errors.ErrorMetaType]interface{}),
			Cause: nil,
		}
	}
	return res
}

// addSpace prefixes each line of the given string with a repeated space pattern.
func addSpace(code, space string, n int) string {
	prefix := strings.Repeat(space, n)
	return prefix + strings.ReplaceAll(code, "\n", "\n"+prefix)
}

func ErrFmt(ei core.ErrorInterface) string {
	return FormatError(ei, nil)
}

// FormatError formats a core.ErrorInterface into a human-readable colored string.
// It handles known error codes (SystemError, GenerationError, ParsingError, parser3 errors)
// and falls back to a generic output for unknown codes.
func FormatError(ei core.ErrorInterface, prev core.ErrorInterface) string {
	if ei == nil {
		return ""
	}

	inner := GetErr(ei)
	meta := ei.GetMeta()
	code := errors.ErrorCodeType(ei.GetCode())
	addFallback := false
	fallbackAdded := false

	var res strings.Builder

	switch code {
	case shared.SystemError:
		res.WriteString(color.Set("[?YW]System error:[?RT]\n"))
		res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))

	case shared.GenerationError:
		res.WriteString(color.Set("[?YW]Generation error:[?RT]\n"))
		res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))

	case errors.ParsingError, parser3.AdapterErrCode:
		fallback := func() {
			res.WriteString(color.Set("[?YW]Parsing error:[?RT]\n"))
			res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))
		}
		if inner != nil {
			if inner.GetCode() == parser3.AdapterErrCode {
				fallbackAdded = true
				FormatError(inner.Unwrap().(core.ErrorInterface), ei)
			} else {
				fallback()
			}
		} else {
			fallback()
		}

	case parser3.ParseErrCode, parser3.GrammarErrCode:
		res.WriteString(color.Set("[?YW]Parser error (2):[?RT]\n"))
		text := ""
		switch v, _ := meta["Code"].(string); v {
		case "UnexpectedToken":
			expected, _ := meta["Expected"].(string)
			got, _ := meta["Got"].(string)
			raw, _ := meta["Raw"].(string)
			text += fmt.Sprintf("Expected '%s', got '%s'", expected, got)
			if raw != "" {
				text += "\n" + addSpace(raw, whereSpace, 1)
			}
		default:
			msg := ei.GetMsg()
			if v != "" {
				text += v + ":"
				if msg != "" {
					text += "\n"
				}
			}
			if msg != "" {
				text += msg
			}
		}
		if text != "" {
			res.WriteString(addSpace(text, errSpace, 1))
		}
	default:
		addFallback = true
	}

	fallback := func() {
		// Generic fallback
		if !addFallback {
			return
		}
		result := res.String()
		if len(result) > 0 && result[len(result)-1] != '\n' {
			res.WriteRune('\n')
		}
		res.WriteString(color.Set("[?BYW]Error:[?YW] " + ei.GetCode() + "[?RT]"))
		msg := ei.GetMsg()
		if msg != "" {
			res.WriteRune('\n')
			res.WriteString(addSpace(msg, errSpace, 1))
		}
	}

	if ei.GetCode() == "RepeatExpr" && prev != nil {
		if prev.GetCode() == "NodeExpr" && prev.GetMsg() == "building node 'file'" &&
			ei.GetMsg() == "expected at least 1 repetition(s), got 0 at idx=0 start=0-1" {
			// while parser in node file and repeats of blocks is not found
			res.WriteString("Do you forget to add block annotation?")
		} else {
			fallback()
		}
	} else {
		fallback()
	}

	if inner != nil {
		if !fallbackAdded {
			res.WriteString("\n")
		}
		result := FormatError(inner, ei)
		res.WriteString(addSpace(result, whereRedSpace, 1))
	}

	return res.String()
}
```

---

# tal/lang/lcproc.go

```go
package lang

import (
	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
)

func NewLexer() *stringParsing.Lexer {
	return stringParsing.NewLexer([]stringParsing.LexerRule{
		{
			Type:    "CODE",
			Pattern: regexp2.MustCompile('"([^"\\]|\\.)*"', 0),
		},
		{
			Type:    "COMMAND",
			Pattern: regexp2.MustCompile('(?m)^\s*--\s*\#(?<cmd>[^\s]+)\s*(?<args>.*?)$', 0),
		},
		{
			Type:    "MAINBLOCK",
			Pattern: regexp2.MustCompile('(?m)^\s*--\s*\@$', 0),
		},
		{
			Type:    "GLOBALBLOCK",
			Pattern: regexp2.MustCompile('(?m)^\s*--\s*\@!$', 0),
		},
		{
			Type:    "BLOCK",
			Pattern: regexp2.MustCompile('(?m)^\s*--\s*\@(?<name>[^\s]+)$', 0),
		},
		{
			Type:    "CODE",
			Pattern: regexp2.MustCompile('(?s).', 0),
		},
	}, &stringParsing.LexerConfig{
		UseBracketBalance: false,
		Brackets:          [][2]string{},
	})
}

func NewParser() *parser3.Adapter {
	p := parser3.NewParser(NewLexer(), parser3.Grammar{
		"file": parser3.Rule{
			Name: "file",
			Expr: parser3.NodeExpr{
				NodeType: "file",
				Expr: parser3.RepeatExpr{
					Expr: parser3.ChoiceExpr{
						Alternatives: []parser3.Expr{
							parser3.NamedExpr{RuleName: "block"},
							parser3.NamedExpr{RuleName: "mainblock"},
							parser3.NamedExpr{RuleName: "globalblock"},
						},
					},
					Min: 1,
				},
			},
		},
		"mainblock": parser3.Rule{
			Name: "mainblock",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "MAINBLOCK"},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"globalblock": parser3.Rule{
			Name: "globalblock",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "GLOBALBLOCK"},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"block": parser3.Rule{
			Name: "block",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "BLOCK"},
						parser3.RepeatExpr{
							Expr: parser3.TokenExpr{TokenType: "COMMAND"},
							Min:  0,
						},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"code": parser3.Rule{
			Name: "code",
			Expr: parser3.NodeExpr{
				NodeType: "code",
				Expr: parser3.RepeatExpr{
					Expr: parser3.TokenExpr{TokenType: "CODE"},
					Min:  0,
				},
			},
		},
	}, "file", nil)
	return &parser3.Adapter{Parser: p}
}
```

---

# tal/lang/process.go

```go
package lang

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/astools"
)

const (
	SysGlobal = "__SYSBLOCK_GLOBALBLOCK"
	SysMain   = "__SYSBLOCK_MAINBLOCK"
)

func Process(code string) (*TalCode, core.ErrorInterface) {
	p := NewParser()
	pn, err := p.Parse(code)
	if err != nil {
		return nil, core.Wrap(errors.ParsingError, err, "%v", p.String())
	}
	return ProcessTalLang(pn)
}

func ProcessTalLang(pn []stringParsing.ParsedNode) (*TalCode, core.ErrorInterface) {
	c := NewTalCode()
	for _, node := range astools.GetChildren(&pn[0]) {
		chs := astools.GetChildren(&node)
		sec := NewTalSection()
		for _, ch := range chs {
			switch ch.Switch {
			case "code":
				sec.Code += ch.Raw
			case "GLOBALBLOCK", "MAINBLOCK":
				sec.Name = "__SYSBLOCK_" + ch.Switch
			case "BLOCK":
				sec.Name = ch.Metadata["name"].(string)
			case "COMMAND":
				sec.Cmds[ch.Metadata["cmd"].(string)] = ch.Metadata["args"].(string)
			default:
				return nil, core.Err(errors.ParsingError, "Unknown: %v", ch.Switch)
			}
		}
		switch sec.Name {
		case SysGlobal:
			c.Global = sec
		case SysMain:
			c.Main = sec
		default:
			c.Blocks[sec.Name] = sec
		}
	}
	return c, nil
}
```

---

# tal/lang/struct.go

```go
package lang

type TalSection struct {
	Cmds map[string]string
	Code string
	Name string
}

func NewTalSection() *TalSection {
	return &TalSection{
		Cmds: make(map[string]string),
		Code: "",
		Name: "",
	}
}

type TalCode struct {
	Global *TalSection
	Main   *TalSection
	Blocks map[string]*TalSection
}

func NewTalCode() *TalCode {
	return &TalCode{
		Global: nil,
		Main:   nil,
		Blocks: make(map[string]*TalSection),
	}
}
```

---

# tal/lua/main.go

```go
package lua

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/pt-main/run/tal/core"
	"github.com/pt-main/tap/go/color"
	lua "github.com/yuin/gopher-lua"
)

type LuaFuncBuilder func(changedFiles, args []string) lua.LGFunction

var GlobalFuncs = map[string]LuaFuncBuilder{}

func NewTalLuaState(changedFiles, args []string) *lua.LState {
	L := lua.NewState()
	L.SetGlobal("changed", L.NewFunction(func(L *lua.LState) int {
		relPaths := makeRelativePaths(changedFiles)
		tbl := L.NewTable()
		for i, p := range relPaths {
			tbl.RawSetInt(i+1, lua.LString(p))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("get_args", L.NewFunction(func(L *lua.LState) int {
		relPaths := makeRelativePaths(args)
		tbl := L.NewTable()
		for i, p := range relPaths {
			tbl.RawSetInt(i+1, lua.LString(p))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("update", L.NewFunction(func(L *lua.LState) int {
		err := core.Update()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 2
		}
		return 1
	}))

	L.SetGlobal("print_colored", L.NewFunction(func(L *lua.LState) int {
		color.PrintColored(L.CheckString(1))
		return 1
	}))

	L.SetGlobal("match_pattern", L.NewFunction(func(L *lua.LState) int {
		pat := L.CheckString(1)
		str := L.CheckString(2)
		ok, err := doublestar.Match(pat, str)
		if err != nil {
			L.Push(lua.LFalse)
			return 1
		}
		L.Push(lua.LBool(ok))
		return 1
	}))

	for name, builder := range GlobalFuncs {
		L.SetGlobal(name, L.NewFunction(builder(changedFiles, args)))
	}
	return L
}

func RegisterLuaFunc(name string, builder LuaFuncBuilder) {
	GlobalFuncs[name] = builder
}

func makeRelativePaths(absPaths []string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		return absPaths
	}
	rel := make([]string, 0, len(absPaths))
	for _, p := range absPaths {
		relPath, err := filepath.Rel(cwd, p)
		if err == nil {
			relPath = filepath.ToSlash(relPath)
			rel = append(rel, relPath)
		} else {
			rel = append(rel, p)
		}
	}
	return rel
}
```

---

# tal/main.go

```go
package tal

import (
	"errors"

	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/lua"
)

const Version = "1.1.0"

func Process(changedFiles, args []string, file string) error {
	ls := lua.NewTalLuaState(changedFiles, args)
	var err_ lccore.ErrorInterface
	processed, err_ := lang.Process(file)
	if err_ != nil {
		return errors.New(lang.ErrFmt(err_))
	}
	generated, err := generation.GenerateCode(processed)
	if err != nil {
		return errors.New(lang.ErrFmt(err))
	}
	return ls.DoString(generated)
}
```

---

# tal/runtime/main.go

```go
package runtime

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/orderedmap"
	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal"
	"github.com/pt-main/run/tal/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
	tap "github.com/pt-main/tap/go"
	"github.com/pt-main/tap/go/color"
)

func CreateCli() *tap.Parser {
	p := tap.NewParser("tal",
		'[?BE]╭─────── [?BRD]Tal[?RT] - Task Lua
[?BE]⎬─ [?RT]Incremental task runner with Lua and file-based dependencies.
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]',
		[]string{"-h", "help"},
		tap.DefaultParserConfig(),
	)

	p.AddCommand("update", UpdateHandler,
		'[?GN]Update .tal.pack[?RT]
Scan the current directory, compute SHA256 hashes for all files,
and store them in a binary .tal.pack file. Run this once before
using 'tal run' to enable incremental builds.
[?BBK]Usage:[?RT]
  [?BBK]tal update[?RT]
[?YW]Example:[?RT]
  [?BBK]tal update[?RT]',
		nil, nil, false)

	p.AddCommand("list", ListHandler,
		'[?GN]List tasks in a Tal file[?RT]
Read a .task.lua file and display all defined tasks (including
global and main blocks). Useful for quickly checking what's
available in a project.
[?BBK]Usage:[?RT]
  [?BBK]tal list <file> [file...][?RT]
[?YW]Example:[?RT]
  [?BBK]tal list main.task.lua[?RT]',
		[]string{"task-lua-file"}, nil, true)

	p.AddCommand("run", RunHandler,
		'[?GN]Run a Tal file[?RT]
Parse and execute the given .task.lua file. All arguments after
the file name are passed directly to the Lua script via get_args()
and are not interpreted by the CLI.
[?BBK]Usage:[?RT]
  [?BBK]tal run <file> [args...] [--deps="dep1;dep2"][?RT]
[?BBK]Flag --deps:[?RT]
  Used to pass dependencies for the --#depends annotation.
  If this flag is not provided, modified files are specified
  as dependencies (requires the presence of .tal.pack).
[?YW]Examples:[?RT]
  [?BBK]tal run main.task.lua[?RT]           # run the script (main block or default)
  [?BBK]tal run main.task.lua build[?RT]     # pass "build" as argument to the script
  [?BBK]tal run main.task.lua test -v[?RT]   # pass arguments to the script',
		[]string{"task-lua-file"}, nil, true)

	p.AddCommand("init", InitHandler,
		'[?GN]Create a default Tal file[?RT]
Generate a minimal main.task.lua file with a main block that
accepts arguments. This is a quick way to start a new project.
[?BBK]Usage:[?RT]
  [?BBK]tal init[?RT]
[?YW]Example:[?RT]
  [?BBK]tal init[?RT]',
		nil, nil, false)

	return p
}

func InitHandler(p *tap.Parser, s []string) error {
	color.PrintlnColored("Update err: %v", core.Update())
	color.PrintlnColored("File creating err: %v", core.Write("main.task.lua", []byte('-- @
if #get_args() > 0 then
    script(get_args()[1]) 
end')))
	return nil
}

func center(s string, width int) string {
	runes := []rune(s)
	n := len(runes)
	if n >= width {
		return s
	}
	left := (width - n) / 2
	right := width - n - left
	return fmt.Sprintf("%*s%s%*s", left, "", s, right, "")
}

func ListHandler(p *tap.Parser, s []string) error {
	for _, fileName := range s {
		file, err := core.OpenF(fileName)
		if err != nil {
			return err
		}
		var err_ lccore.ErrorInterface
		parsed, err_ := lang.Process(file)
		if err_ != nil {
			return errors.New(lang.ErrFmt(err_))
		}
		res := []string{"[?GN]╭─────── [?RT][[?YW]" +
			center(fileName, 20) + "[?RT]] Scripts"}
		templ := "[?GN]│  [?RT]%3d: [?BGN]%v"
		idx := 0
		if parsed.Global != nil {
			idx += 1
			res = append(res, fmt.Sprintf(templ, idx, "Global"))
		}
		for script := range parsed.Blocks {
			res = append(res, fmt.Sprintf(templ, idx, script))
			idx += 1
		}
		if parsed.Main != nil {
			res = append(res, fmt.Sprintf(templ, idx, "Main"))
		}
		res = append(res, "[?GN]╰───────[?RT]")
		color.PrintlnColored(strings.Join(res, "\n"))
	}
	return nil
}

func UpdateHandler(p *tap.Parser, s []string) error {
	return core.Update()
}

func RunHandler(p *tap.Parser, s []string) (err error) {
	ch := []string{}
	if deps, hasDeps := p.Flags["deps"]; hasDeps {
		ch = strings.Split(deps, ";")
	} else {
		ch, err = GetChanges()
		if err != nil {
			return
		}
	}
	args := []string{}
	skippedName := false
	for _, arg := range p.RawArgs[1:] {
		if arg == s[0] && !skippedName {
			skippedName = true
		} else {
			args = append(args, arg)
		}
	}
	file, err := core.OpenF(s[0])
	return tal.Process(ch, args, file)
}

func GetSavedFile() (*orderedmap.OrderedMap, error) {
	file, err := core.Open(shared.TalFile)
	if err != nil {
		return nil, err
	}
	return core.PackCoreAsState(file)
}

func GetChanges() ([]string, error) {
	w, err := GetSavedFile()
	if err != nil {
		return nil, err
	}
	return core.Changes(w, ".")
}
```

---

# tal/shared/config.go

```go
package shared

const TalFile = ".tal.pack"
```

---

# tal/shared/errors.go

```go
package shared

import "github.com/pt-main/lc/public/errors"

const (
	GenerationError errors.ErrorCodeType = "GENERATION"
	SystemError     errors.ErrorCodeType = "SYSTEM"
)
```

---

# tal/test/main.go

```go
package main

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
)

func main() {
	fmt.Println("start")
	code := '
-- @build
-- #depends main.go test.go
print("start building...")

-- @
run("build")'
	proc, err := lang.Process(code)
	fmt.Println(core.GetRealError(err), proc)

	res, err := generation.GenerateCode(proc)
	fmt.Println(core.GetRealError(err), res)
}
```

---

# tal/test/main.task.lua

```lua
-- @test
print("test")
run_cli("--lm -r test")
run_cli("tal run main.task.lua test2")

-- @test2
print("test2")
script("test3")
script("test4")

-- @test3
-- #depends main.go
print("1!")

-- @test4
-- #depends merged.md
print("2!")

-- @
if #get_args() > 0 then
    script(get_args()[1]) 
end

update()
```

---

# tal/test/simple.task.lua

```lua
-- @
print("Working")
```

---

