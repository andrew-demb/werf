<!-- THIS FILE IS AUTOGENERATED BY werf docs COMMAND! DO NOT EDIT! -->

**werf** — Open Source CLI-утилита, написанная на Go, предназначенная для упрощения и ускорения доставки вашего приложения.

Вам достаточно описать конфигурацию приложения, правила сборки и развертывания в Kubernetes, в Git-репозитории, едином источнике правды. Проще говоря, это то, что сегодня называется GitOps.

* Собирает Docker-образы, как используя Dockerfile, так и альтернативный сборщик с собственным синтаксисом, основная задача которого — сокращение времени инкрементальной сборки на основе истории Git.
* Поддерживает множество схем тегирования.
* Выкатывает приложение в Kubernetes, используя Helm-совместимый формат чартов с удобными настройками, улучшенным механизмом отслеживания процесса выката, обнаружения ошибок и выводом логов.
* Очищает Docker registry от неиспользуемых образов.

werf — не CI/CD-система, а инструмент для построения пайплайнов, который может использоваться в любой CI/CD-системе. Мы считаем инструменты такого рода новым поколением высокоуровневых инструментов CI/CD.
