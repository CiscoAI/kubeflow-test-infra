#!/bin/bash
#usage: ./masterkubeup.sh [masterIPaddress]
if [ -z "$1" ]
then
	ipaddress="0.0.0.1"
else
	ipaddress="$1"
fi
sudo systemctl enable docker
sudo systemctl start docker
sudo systemctl enable kubelet
sudo systemctl start kubelet
sudo systemctl daemon-reload
sudo systemctl restart kubelet
sudo kubeadm config images pull
sudo kubeadm reset --force
sudo rm -rf $HOME/.kube
sudo swapoff -a
# Needs ignore preflight errors because of docker version 18.09.0 from nvidia-docker which is unsupported in kubeadm
sudo kubeadm init --apiserver-advertise-address=$ipaddress --pod-network-cidr=192.168.0.0/16
rm -R $HOME/.kube
mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=$HOME/.kube/config
kubectl apply -f $HOME/dex-istio-kf/calico/
#kubectl apply -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
# Uncomment next line for single node cluster
kubectl taint nodes --all node-role.kubernetes.io/master-
