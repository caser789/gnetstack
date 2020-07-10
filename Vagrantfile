Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.define "httpserver" do |node|
    node.vm.hostname = "httpserver"
    node.vm.network "private_network", ip: "192.168.100.10"
    node.vm.provider "virtualbox" do |v|
        v.memory = 512
        v.cpus = 1
    end
  end
end
