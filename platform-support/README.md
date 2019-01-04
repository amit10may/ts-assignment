### Solution for Platform Support assignments
- The main bash script test.sh contains the code to accept input and count tokens
- The test1.exp and test2.exp scripts contains "expect" commands to test the test.sh script
- Ansible script copy_script.yml contains code to transfer the test.sh to remote server
- Ansible script cope_autoscript.yml contains code to transfer the testing script to remote server (test1.exp)
- Ansible script run_script.yml contains code to run the autoexpect script test1.exp remotely.
- test.sh cannot be run remotely since it needs user interaction. So used expect approch
- Screenshots are contained the solution document (docx)

### Notes
- ansible.cfg needed small change to disable host checking
- hosts / inventory file not uploaded for security reasons. Below is the snippet from the same
- [dev]
- x.x.x.x ansible_connection=ssh ansible_user=user ansible_ssh_pass=secretpwd ansible_become_pass=secretsudopwd

### Pending
- Writing the rsyslog filter
- Writing the Ansible Role to test the script
