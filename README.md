# fm_mindmap_app

fm_mindmap を Wails を使ってアプリケーション化するリポジトリ

## 概要

frontend 下に git submodule として fm-mindmap (https://github.com/rerofumi/fm-mindmap)を展開し、Wails プロジェクトとしてビルドします。
ビルド後は 1つの実行ファイルとしてアプリ化されますので、利用しやすくなります。

LLM 機能の利用には openrouter キーの設定が必要ですが、それは .env に記載して、exe と同じディレクトリに配置します。

## 環境構築

ビルド環境の構築、ビルドランナーとして mise-en-place (https://mise.jdx.dev/) を使用します。
開発ツールとして以下を衣装しますが、これらはすべて mise が準備します。
- golang
- wails
- node.js

ビルドを進めるには mise のみインストールしてください。
Windows 環境では `winget install jdx.mise` でインストールできます。

この git リポジトリの他に submodule を clone する必要があります
submodule clone 置閏

## ビルド

最初に `mise install` で必要なツールをインストールします。

次に `mise run setup` で wails の準備をします。

`mise run dev` で開発環境の実行、`mise run build` でアプリのビルドです。

## 設定ファイル(.env)

アプリでは LLM 機能を利用しますが、そのために openrouter のAPIキーと利用クレジットが必要です。
dev 実行時はこのリポジトリ直下に .env ファイルを作成してください。ビルドアプリを実行するときは、その実行ファイルと同じディレクトリに .env ファイルを置いて下さい。
