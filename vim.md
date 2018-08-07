## vim Install and Update
    git clone https://github.com/vim/vim.git

    * update to the latest version with:
    cd vim
    git pull

    cd src
    make distclean  # if you build Vim before
    make
    sudo make install

## vim的命令
    vim filename            打开或新建文件，并将光标置于第一行首
    vim +n filename         打开文件，并将光标置于第n行首
    vim + filename          打开文件，并将光标置于最后一行首
    vim +/pattern filename  打开文件，并将光标置于第一个与pattern匹配的串处
    vim -r filename         在上次正用vim编辑时发生系统崩溃，恢复filename
    vim filename….filename  打开多个文件，依次编辑

## Vim配置：
    all                     列出所有选项设置情况
    term                    设置终端类型
    ignorance               在搜索中忽略大小写
    list                    显示制表位(Ctrl+I)和行尾标志（$)
    number                  显示行号
    report                  显示由面向行的命令修改过的数目
    terse                   显示简短的警告信息
    warn                    在转到别的文件时若没保存当前文件则显示NO write信息
    nomagic                 允许在搜索模式中，使用前面不带“\”的特殊字符
    nowrapscan              禁止vi在搜索到达文件两端时，又从另一端开始
    mesg                    允许vi显示其他用户用write写到自己终端上的信息
    :set number / set nonumber 显示/不显示行号
    :set ruler /set noruler    显示/不显示标尺
    :set hlsearch              高亮显示查找到的单词
    :set nohlsearch            关闭高亮显示
    :syntax on                 语法高亮
    :set tabstop=4             设置tab大小
    :set softtabstop=4
    :set autoindent            自动缩进
    :set cindent               C语言格式里面的自动缩进

## 移动光标：
    k   上   nk     向上移动n行
    j   下   nj     向下移动n行
    h   左   nh     向左移动n列
    l   右   nl     向右移动n列
    space           光标右移一个字符
    Backspace       光标左移一个字符
    Enter           光标下移一行
    w或W            光标右移一个字至字首
    b或B 			光标左移一个字至字首
    e或E 			光标右移一个字至字尾
    ) 			光标移至句尾
    ( 			光标移至句首
    }			光标移至段落开头
    {			光标移至段落结尾
    n$			光标移至第n行尾
    H 			光标移至屏幕顶行
    M 			光标移至屏幕中间行
    L 			光标移至屏幕最后行
    0			（注意是数字零）光标移至当前行首
    ^			移动光标到行首第一个非空字符上去
    $			光标移至当前行尾
    gg          可以移到第一行
    G           移到最后一行
    f<a>        移动光标到当前行的字符a上
    F           相反
    %           移动到与制匹配的括号上去（），{}，[]，<>等。
    nG          移动到第n行上

## 屏幕翻滚类命令
    Ctrl+u		向文件首翻半屏
    Ctrl+d		向文件尾翻半屏
    Ctrl+f		向文件尾翻一屏
    Ctrl＋b		向文件首翻一屏
    nz			将第n行滚至屏幕顶部，不指定n时将当前行滚至屏幕顶部

## 插入文本类命令
    i			在光标前
    I			在当前行首
    a			光标后
    A			在当前行尾
    o			在当前行之下新开一行
    O			在当前行之上新开一行
    r			替换当前字符
    R			替换当前字符及其后的字符，直至按ESC键
    s			从当前光标位置处开始，以输入的文本替代指定数目的字符
    S			删除指定数目的行，并以所输入文本代替之
    ncw或nCW	修改指定数目的字
    nCC			修改指定数目的行

## 删除命令
    x或X		删除一个字符，x删除光标后的，而X删除光标前的
    dw			删除一个单词
    dnw         删除n个单词
    dne         也可，只是删除到单词尾
    d0			删至行首
    d$			删至行尾
    dd			删除一行
    ndd			删除当前行及其后n-1行
    dnl			向右删除n个字母
    dnh			向左删除n个字母
    dnj			向下删除n行
    dnk			向上删除n行
    cnw[word]	将n个word改变为word
    C$			改变到行尾
    cc			改变整行
    shift+j     删除行尾的换行符，下一行接上来了.
    :g/^#/d	    删除所有以「#」打头的行
    :g/^$/d 	删除所有空行
    :g/^\s*$/d  删除只有空白的行

## 复制粘贴
    p			粘贴用x或d删除的文本
    ynw			复制n个单词
    yy			复制一行
    ynl			复制n个字符
    y$			复制当前光标至行尾处
    nyy			拷贝n行

## 撤销
    u			Undo
    shif+u(U)	撤销对该行的所有操作。
    Ctrl+r		Redo

## 搜索及替换命令 			
    /pattern	    从光标开始处向文件尾搜索pattern
    ?pattern	    从光标开始处向文件首搜索pattern
    n			    在同一方向重复上一次搜索命令
    N			    在反方向上重复上一次搜索命令
    cw【newword】	替换为newword
    n			    继续查找
    .		    	执行替换
    :s/xxx/yyy/ 	将当前行中的第一个「xxx」替换为「yyy」
    :s/xxx/yyy/g	将当前行中的所有「xxx」替换为「yyy」
    :s/xxx/yyy/gc	同上，但每次替换都会询问
    :%s/xxx/yyy/g	将整个文件中的所有「xxx」替换为「yyy」

    * 正则用法
        \(and\) 表示占位符,\1获取匹配
        \<word\> 表示匹配单个单词

        /\(a\+\)[^a]\+\1        查找开头和结尾处a的个数相同的字符串，如aabbbaa，aaacccaaa，但是不匹配abbbaa
        :s/\<four\>/4/g         将所有的four替换成4，但是fourteen中的four不替换
        :%s/^\([^\ ]\)/## \1/   行首非空格添加'## '

    * 函数式
        :s/替换字符串/\=函数
        在函数式中可以使用 submatch(1)、submatch(2) 等来引用 \1、\2 等的内容，而submatch(0)可以引用匹配的整个内容。
        
        :%s/\<id\>/\=line(".") 将各行的id替换为行号
        :%s/^\<\w\+\>/\=(line(".")-10) .".". submatch(1) 将每行开头的单词替换为 (行号-10).单词 的格式，如第11行的 word 替换成 1. word

    :n1,n2 s/p1/p2/g            将第n1至n2行中所有p1均用p2替代
    :g/p1/s//p2/g               将文件中所有p1均用p2替换
    :1,$ s/p1/p2/g              在全文中将p1替换为p2

## 文本替换命令的格式为
    在文件及目录中查找字符串的方法
    :vimgrep /{pattern}/[g][j] {file}
    :vimgrep /hello world/g **/*.*
    查找的结果可以用 :copen 命令打开 quickfix 列表查看
    没有参数 g 的话,则行只查找一次关键字.反之会查找所有的关键字
    没有参数 j 的话,查找后, VIM 会跳转至第一个关键字所在的文件.反之,只更新结果列表(quickfix)
    :cc    是在转到当前查找到的位置
    :cn    转到下一个位置
    :cp    转到前一个位置

## 书签
    m[a-z]			在文中做标记，标记号可为a-z的26个字母
    `a		    	移动到标记a处

## 行方式命令
    :n1,n2 co n3		将n1行到n2行之间的内容拷贝到第n3行下
    :n1,n2 m n3			将n1行到n2行之间的内容移至到第n3行下
    :n1,n2 d 			将n1行到n2行之间的内容删除
    :n1,n2 w!command	将文件中n1行至n2行的内容作为command的输入并执行之
                        若不指定n1，n2，则表示将整个文件内容作为command的输入

## 窗口操作
    <C-w>s 或 :split    横向分屏
    <C-w>v 或 :vsplit   垂直分屏
    :split file.c       为另一个文件file.c分隔窗口
    :nsplit file.c      为另一个文件file.c分隔窗口，并指定其行数
    ctrl＋w             在窗口中切换
    :close              关闭当前窗口


## 文件及其他
    :q                   退出vi
    :q!                  不保存文件并退出vi
    :e filename          打开文件filename进行编辑
    :e!                  放弃修改文件内容，重新载入该文件编辑
    :w                   保存当前文件
    :wq                  存盘退出
    :ZZ                  保存当前文档并退出VIM
    :!command            执行shell命令command
    :r!command           将命令command的输出结果放到当前行
    :n1,n2 write temp.c  将本文件中的n1,到n2行写入temp.c这个文件中去
    :read file.c         将文件file.c的内容插入到当前光标所在的下面
    :0read file.c        将文件file.c的内容插入到当前文件的开始处(第0行）
    :nread file.c        将文件file.c的内容插入到当前文件的第n行后面
    :read !cmd           将外部命令cmd的输出插如到当前光标所在的下面
    :f newFile           重命名文件

    文件编码批量转换
    :args */**/*.txt
    :argdo set fenc=utf-8 | update

## 帮助
    :help               查看帮助文档，在这之中，按CTRL+] 进入超连接，按CTRL＋O 返回。
    :help subject       看某一主题的帮助，ZZ 退出帮助

## QuickFix 窗口
    :cw                 如果有编译错误信息就会打开 QuickFix 窗口
    :ccl                关闭 QuickFix 窗口

## 其它常用：
    :au GUIEnter * simalt ~x   启动时最大化

## 查看命令
    :vmap

## .vimrc 自定义命令
    The leader is ,, so whenever you see <leader> it means ,.

    ,cd      切换工作目录CWD到当前打开buffer的目录
    ,w :w!   写入

    tab
    ,tn :tabnew
    ,tc :tabclose
    ,to :tabonly
    ,tm :tabmove
    ,te :tabedit

## tab
    :newtab   新建标签
    :gt       转至下一个标签
    :gT       转到上一个标签
    
## 简单代码跳转
    `.        跳转到最后编辑行
    gd        到局部变量的定义
    gD        到全局变量的定义
    
    *, #      可对光标处的词向前/向后做全词搜索，g*, g# 做相对应的非全词匹配搜索
    [[, ]]    可跳到当前代码块起始或结束的大括号处。
    %         可在配对的括号、块首尾之前跳转。

    Ctrl-O    在历史记录中后退
    Ctrl-I    在历史记录中前进。

## visual 模式：
    v         进入visual 模式
    V         进入行的visual 模式
    J         把所有的行连接起来(变成一行)
    < 或 >    左右缩进
    =         自动给缩进
    ctrl+v    进如块操作模式用o和O改变选择的边的大小
    
    在所有行插入相同的内容：
    将光标移到开始插入的位置，按CTRL+V进入VISUAL模式，选择好模块后按I（shift+i)，输入内容，按[ESC]完成

    在所有被选择的行前加上点东西：
    ^ 到行首
    <C-v> 开始快操作
    <C-d> 向下移动或hjkl
    I 输入内容, 按[Esc]键为每一行生效

    在所有被选择的行后加上点东西：
    <C-v>
    选中相关行(hjkl)
    $ 到行最后
    A 输入内容，按Esc

## 宏录制
    q[a-z]：开始记录但前开始的操作为宏，名称可为【a-z】，然后用q终止录制宏
    reg：显示当前定义的所有的宏，用@[a-z]来在当前光标处执行宏[a-z]

    qa把你的操作记录在寄存器a中
    于是@a会replay被录制的宏
    @@是一个快捷键用来replay最新录制的宏

    在一个只有一行且这一行只有“1”的文本中，键入如下命令：
    qa 开始录制
    Yp 复制行
    <C-a> 增加1
    q 停止录制
    @a 在1下面写下 2
    @@ 在2 正面写下3
    现在做 100@@ 会创建新的100行，并把数据增加到 103.

## 实现插入递增数字
    :let i=0 | g/toPlaced/s//\=i/ |let i=i+1
    这条命令由三部分组成:
    let i=1 和 let i=i+1 构成了一个变量递增的循环。
    g命令用于全局查找一个字符串，并对此字符串施加一个命令。例如:
    g/\(12\)3/s//\14/  #查找123，并将3替换成4

    将匹配的数字加2
    例如 No.1
         No.2
         :%s/\(No.\)\(\d\+\)/\=submatch(1).".".(submatch(2)-2)
    结果 No.3
         No.4
    submatch(1) 匹配 \(No.\)
    . 连接符
    "." 转义 . 符号

## 关键字补全
    即简单地补全到文档中已有的词，快捷键为 Ctrl-N 或 Ctrl-P。

## 浏览目录
    :E
    :He 在下边分屏浏览目录
    :He! 在上边分屏浏览目录
    :Ve 在左边分屏浏览目录

    - 到上级目录
    D 删除文件
    R 重命名
    s 对文件排序
    x 执行文件

## 缓冲区
    :ls
    :buffer 3 切换文件到3
    :bnext 简写 :bn
    :bprevious
    :bfirst
    :blast
