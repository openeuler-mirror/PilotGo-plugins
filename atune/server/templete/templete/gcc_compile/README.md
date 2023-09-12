1. Prepare the environment  
sh prepare.sh
2. Start to tuning  
atune-adm tuning --project gcc_compile --detail gcc_compile_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project gcc_compile
