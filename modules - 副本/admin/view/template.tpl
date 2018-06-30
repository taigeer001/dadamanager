<style>
    #mobans .card .body{
    padding:0px;
    }
    #mobans .card .body img{
        width: 100%;
    }
</style>
<div class="block-header">
<h2>ADVANCED FORM ELEMENTS</h2>
</div>
<div class="row" id="mobans">
    {{range .}}
    <div class="col-lg-2 col-md-2 col-sm-3 col-xs-6">
        <div class="card">
            <div class="header bg-orange">
                <h2>
                    {{.Author}}@{{.Name}} <small>{{.Description}}</small>
                </h2>
                <ul class="header-dropdown m-r--5">
                    <li class="dropdown">
                        <a href="javascript:void(0);" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                            <i class="material-icons">more_vert</i>
                        </a>
                        <ul class="dropdown-menu pull-right">
                            <li><a href="javascript:void(0);" class=" waves-effect waves-block">Action</a></li>
                            <li><a href="javascript:void(0);" class=" waves-effect waves-block">Another action</a></li>
                            <li><a href="javascript:void(0);" class=" waves-effect waves-block">Something else here</a></li>
                        </ul>
                    </li>
                </ul>
            </div>
            <div class="body">
                <img src="/icon/{{.Dir}}/icon.png" />
            </div>
        </div>
    </div>
    {{end}}
</div>