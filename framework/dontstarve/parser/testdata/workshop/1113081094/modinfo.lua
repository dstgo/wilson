name = "功能合集"
description = 
[[添加冰箱、盐箱可设置永久保鲜、返鲜；
添加物品叠加可以修改，最高999；
添加帐篷、小木棚使用次数可修改；
添加小猪包、坎普斯背包、大理石甲可修改移动速度（移除减速会改变配方）；
添加快速采集功能；
添加宝石制作功能（绿宝石、橙宝石、黄宝石）；
添加木栅栏和木门无法攻击摧毁功能（能锤掉）；
添加围墙护甲（100%护甲其实是99.99%，T键的一键必杀可以打掉，正常情况打不掉）；
添加草蜥蜴转化（没移除，官方已支持开局设置不转化，未测试开档后改还行不行）；
添加热能石不掉耐久；
添加切斯特100%防御（防止小切被狗一波带走，说多了都是泪）；
添加锤子清理物品功能（右键给予，可以吃很多东西，可以自己去modmain里面添加，如果不会的话，留言我有空加）；
添加护甲、武器、镐斧、毛皮铺盖卷耐久度修改；
以上涉及到数值改动的，都可以自己修改，开档的时候自己打开MOD配置那个文件就可以改成想要的数值。
]]
author = "zzz"
version = "1.81"

forumthread = ""

api_version = 10
dst_compatible = true
reign_of_giants_compatible = false

all_clients_require_mod = true
client_only_mod = false

icon_atlas = "modicon.xml"
icon = "modicon.tex"

configuration_options =
{

	{
	  name = "icebox",
	  label = "冰箱设置",
	  hover = "icebox setting",
	   options = 
        {
	       	{description = "默认", data = .5, hover = "default"},
			{description = "永久保鲜", data = 0, hover = "keep fresh forever"},
			{description = "返鲜", data = -1, hover = "return to freshness"},
        },
	    default = .5,
	
	},

	{
	  name = "saltbox",
	  label = "盐箱设置",
	  hover = "saltbox setting",
	   options = 
        {
	       	{description = "默认", data = .25, hover = "default"},
			{description = "永久保鲜", data = 0, hover = "keep fresh forever"},
			{description = "返鲜", data = -1, hover = "return to freshness"},
        },
	    default = .25,
	
	},
	
	{
	  name = "size",
	  label = "物品叠加",
	  hover = "the number of stacked",
	   options = 
        {		
			{description = "10", data = 10},
			{description = "20", data = 20},
			{description = "40", data = 40},
			{description = "60", data = 60},
			{description = "80", data = 80},
			{description = "99", data = 99},
			{description = "120", data = 120},
			{description = "160", data = 160},
			{description = "200", data = 200},
			{description = "300", data = 300},
			{description = "500", data = 500},
			{description = "999", data = 999},
     	 },
	    default = 40,
	
	},
	
    {
       name = "tent",
       label = "帐篷耐久",
	   hover = "tent to used",
        options = 
        {
			 {description = "3", data = 3},
			 {description = "6", data = 6},
			 {description = "9", data = 9},
			 {description = "15", data = 15},
             {description = "30", data = 30},
			 {description = "45", data = 45},
			 {description = "60", data = 60},
             {description = "75", data = 75},
             {description = "90", data = 90},
        },
        default = 15,
    },
	
	{
	   name = "siestacanopy",
       label = "小木棚耐久",
	   hover = "siesta canopy to used",
        options = 
        {
			 {description = "3", data = 3},
			 {description = "6", data = 6},
			 {description = "9", data = 9},
			 {description = "15", data = 15},
             {description = "30", data = 30},
			 {description = "45", data = 45},
			 {description = "60", data = 60},
             {description = "75", data = 75},
             {description = "90", data = 90},
        },  
        default = 15,
	},

    {
	   name = "piggybackspeed",
	   label = "小猪包移速",
	   hover = "piggy backpack moving speed（移除减速会改变配方/Removing the deceleration will change the recipe)",
		options =	
		{
			 {description = "默认", data = 0.9, hover = "default"},
			 {description = "移除减速", data = 1, hover = "remove deceleration"},
			 {description = "1.1倍移速", data = 1.1, hover = "1.1 times speed"},
		     {description = "1.2倍移速", data = 1.2, hover = "1.2 times speed"},
		},
		default = 0.9,
	},
	
	{
	   name = "krampussack",
	   label = "坎普斯背包移速",
	   hover = "krampus backpack moving speed",
		options =	
		{
		     {description = "默认", data = 1, hover = "default"},
			 {description = "1.1倍移速", data = 1.1, hover = "1.1 times speed"},
			 {description = "1.2倍移速", data = 1.2, hover = "1.2 times speed"},
		     {description = "1.3倍移速", data = 1.3, hover = "1.3 times speed"},	 
		},
		default = 1,
	},
	
   	{
	   name = "armormarble",
	   label = "大理石盔甲移速",
	   hover = "armor marble moving speed（移除减速会改变配方/Removing the deceleration will change the recipe)",
		options =	
		{
		     {description = "默认", data = 0.7, hover = "default"},
			 {description = "移除减速", data = 1, hover = "remove deceleration"},
			 {description = "1.1倍移速", data = 1.1, hover = "1.1 times speed"},
		     {description = "1.2倍移速", data = 1.2, hover = "1.2 times speed"},
		},
		default = 0.7,
	},
	
	{
	   name = "pick",
	   label = "快速工作",
	   hover = "quick work",
		options =	
		{
			 {description = "关闭", data = false, hover = "close"},
		     {description = "开启", data = true, hover = "open"},
		},
		default = false,
	},
	
	{
       name = "consume",
       label = "宝石制作",
	   hover = "the difficulty of making gems",
        options = 
        {
			 {description = "关闭", data = "no", hover = "close"},
		     {description = "简单", data = "easy", hover = "easy"},
    	     {description = "正常", data = "normal", hover = "normal"},
		     {description = "困难", data = "hard", hover = "hard"},
        },
		default = "no",
    },
	
	{
       name = "fence_second",
       label = "木栅栏、木门能否被攻击摧毁",
	   hover = "fence and door can be destroyed by attack?",
        options = 
        {
			 {description = "可以", data = true, hover = "yes"},
		     {description = "不能", data = false, hover = "no"},

        },
		default = true,
    },
	
	{
       name = "wall_armor",
       label = "围墙护甲",
	   hover = "wall armor",
        options = 
        {
			 {description = "0%", data = 0},
			 {description = "25%", data = 0.25},
		     {description = "50%", data = 0.5},
		     {description = "75%", data = 0.75},
		     {description = "100%", data = 0.9999},
        },
		default = 0,
    },
	
	{
       name = "no_grassgekko",
       label = "草蜥蜴转化",
	   hover = "no grassgekko",
        options = 
        {
		     {description = "转化", data = false, hover = "false"},
			 {description = "不转化", data = true, hover = "true"},
        },
		default = false,
    },
	
	{
       name = "heatrock_u",
       label = "热能石不掉耐久",
	   hover = "heat rock whether out of durable",
        options = 
        {
		     {description = "关闭", data = false, hover = "false"},
			 {description = "开启", data = true, hover = "true"},
        },
		default = false,
    },

	{
       name = "chester_a",
       label = "切斯特添加防御",
	   hover = "chester add armor",
        options = 
        {
		     {description = "关闭", data = false, hover = "false"},
			 {description = "开启", data = true, hover = "true"},
        },
		default = false,
    },
	
	{
       name = "clean_t",
       label = "锤子添加清理功能",
	   hover = "hammer Add cleaning function",
        options = 
        {
		     {description = "关闭", data = false, hover = "false"},
			 {description = "开启", data = true, hover = "true"},
        },
		default = false,
    },
	
	{
       name = "armour_health",
       label = "护甲耐久",
	   hover = "Armor durability",
        options = 
        {
		     {description = "默认", data = 1, hover = "default"},
			 {description = "2倍", data = 2, hover = "2X"},
			 {description = "3倍", data = 3, hover = "3X"},
		     {description = "4倍", data = 4, hover = "4X"},
             {description = "5倍", data = 5, hover = "5X"},
             {description = "10倍", data = 10, hover = "10X"},
        },
		default = 1,
    },
	
	{
       name = "arms_uses",
       label = "武器耐久",
	   hover = "Weapon durability",
        options = 
        {
		     {description = "默认", data = 1, hover = "default"},
			 {description = "2倍", data = 2, hover = "2X"},
			 {description = "3倍", data = 3, hover = "3X"},
		     {description = "4倍", data = 4, hover = "4X"},
             {description = "5倍", data = 5, hover = "5X"},
             {description = "10倍", data = 10, hover = "10X"},
        },
		default = 1,
    },

	{
       name = "pickaxe_uses",
       label = "镐斧耐久",
	   hover = "Pickaxe durability",
        options = 
        {
		     {description = "默认", data = 1, hover = "default"},
			 {description = "2倍", data = 2, hover = "2X"},
			 {description = "3倍", data = 3, hover = "3X"},
		     {description = "4倍", data = 4, hover = "4X"},
             {description = "5倍", data = 5, hover = "5X"},
             {description = "10倍", data = 10, hover = "10X"},
        },
		default = 1,
    },

	{
       name = "bedroll_uses",
       label = "毛皮铺盖卷耐久",
	   hover = "bedroll furry durability",
        options = 
        {
		     {description = "默认", data = 1, hover = "default"},
			 {description = "2倍", data = 2, hover = "2X"},
			 {description = "3倍", data = 3, hover = "3X"},
		     {description = "4倍", data = 4, hover = "4X"},
             {description = "5倍", data = 5, hover = "5X"},
             {description = "10倍", data = 10, hover = "10X"},
        },
		default = 1,
    },
	
	
}
	
	
	
	
	
	
	
	
	
	
	




