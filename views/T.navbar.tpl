{{define "navbar"}}

		  	<a href="/" class="navbar-brand">个人博客</a>
		  	<ul class="nav navbar-nav">
	            <li {{if .IsHome}} class="active"{{end}}><a href="/">主页</a></li>
	            <li {{if .IsCatalog}} class="active"{{end}}><a href="/catalog">分类</a></li>
	            <li {{if .IsTopic}} class="active"{{end}}><a href="/topic">文章</a></li>
	          </ul>
		    <a href="/login" class="btn btn-default navbar-btn pull-right">登录</a>

{{end}}