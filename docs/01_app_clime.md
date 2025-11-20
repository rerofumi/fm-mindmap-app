vite+typescript で作ったブラウザアプリを wails v2 でビルドしてアプリにする
frontend ディレクトリに fm-mindmap というブラウザアプリが展開されている、これがターゲット
単純にアプリとしてパッケージするのみで frontend 内のコードは変更しない
開発環境として mise を利用する。
Node.js, Go, Wails は mise tools で管理し、起動やビルドは mise task を介して行う。
