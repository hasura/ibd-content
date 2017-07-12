curl https://raw.githubusercontent.com/hasura/ibd-content/master/3.5-supervise/ibd-local.service > ibd-local.service && \
    sudo mv ibd-local.service /etc/systemd/system/ && \
    sudo systemctl enable ibd-local.service
