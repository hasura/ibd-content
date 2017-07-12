curl https://raw.githubusercontent.com/hasura/ibd-content/master/3.5-supervise/ibd-coreos.service > ibd-coreos.service && \
    sudo mv ibd-coreos.service /etc/systemd/system/ && \
    sudo systemctl enable ibd-coreos.service
