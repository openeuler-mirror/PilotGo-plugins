1. Prepare the environment  
sh prepare.sh enwik8.zip
2. Start to tuning  
atune-adm tuning --project compress_Except_example --detail compress_Except_example_client.yaml
3. Restore the environment  
atune-adm tuning --restore --project compress_Except_example
4. Divide ip  
Use "-" to divide inter group ip and "," to divide intra group ip in /etc/atuned/atuned.cnf
5. Use Except  
Use "Except" in server.yaml file to declare which parameters cannot adjust the "object", and ips are separated by ","