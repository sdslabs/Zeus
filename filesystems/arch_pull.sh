docker pull archlinux:base
docker create --name temp_arch archlinux:base
docker export temp_arch > arch_fs.tar
mkdir arch
sudo tar xf arch_fs.tar -C arch
docker rm temp_arch