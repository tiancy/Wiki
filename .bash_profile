export PATH=$PATH:/usr/local/mongodb/bin

export PATH=$PATH:/usr/local/mysql/bin
alias mysql-start='sudo /usr/local/mysql/support-files/mysql.server start'
alias mysql-stop='sudo /usr/local/mysql/support-files/mysql.server stop'

export JAVA_HOME=$(/usr/libexec/java_home)
#export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_111.jdk/Contents/Home

export PATH=$PATH:/usr/local/apache-maven-3.5.3/bin

export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin

export PATH=$PATH:/usr/local/protoc-3.5.0-osx-x86_64/bin

# set terminal color
#export PS1="\[\033[32m\]\u@\h\[\033[m\]:\w\$ "
export PS1="\[\e[30;1m\]\u@\h:\w\$\[\e[m\] "
alias ls='ls -G'

#alias go='https_proxy=127.0.0.1:1087 go'
#alias git='https_proxy=127.0.0.1:1087 git'
