1. Prepare the environment  
sh prepare.sh enwik8.zip
2. Start to tuning  
atune-adm tuning --project compress --detail compress_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project compress
