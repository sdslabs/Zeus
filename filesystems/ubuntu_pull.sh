docker pull ubuntu:latest
docker create --name temp_ubuntu ubuntu:latest
docker export temp_ubuntu > ubuntu_fs.tar
mkdir ubuntu
sudo tar xf ubuntu_fs.tar -C ubuntu
rm ubuntu_fs.tar
docker rm temp_ubuntu