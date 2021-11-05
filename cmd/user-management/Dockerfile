FROM jrei/systemd-centos:7

COPY UserManagement.viper.yaml /etc/UserManagement/UserManagement.viper.yaml
COPY user-management /usr/local/bin/user-management

CMD ["user-management", "run"]
