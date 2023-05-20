# golang_udemy1
Learning Golang with Udemy video materials.

## Learning content
- Golang Basics to Applications
- Standard Package
- Third Package
- Simple CRUD system (ToDo application) development
- Deploy to Heroku
- Introduction to Concurrency Processing
---

2023/5/16
DBをSQLite3からPostgresに変更し、Herokuにデプロイ後
一瞬だけDynoとPostgresのPlanMiniに課金<br>
一通りCRUDの期待通りの挙動を確認後DBのバックアップを取り
`pg_restore`コマンドにて #Supabase へ移行。

5/17 ローカルで再度SQLite3で動く形に`git reset`
`master`→`main`ブランチにて続行

以後、アプリを動く形でどこに公開するか要検討<br>
Supabaseを継続利用して未知の無償サービスと組み合わせるか、<br>
ある程度継続課金してAWSかGCP（Cloud run）、或いはAzure？<br>
どうせ小規模なのでNoSQLに対応させFirebase（Firestore）でもいいかも？

---
