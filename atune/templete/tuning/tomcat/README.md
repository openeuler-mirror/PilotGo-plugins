1. Prepare the environment (with tomcat_root as your tomcat folder)  
sh prepare.sh tomcat_root
2. Start to tuning  
atune-adm tuning --project tomcat --detail tomcat.yaml
3. Restore the environment  
atune-adm tuning --restore --project tomcat
