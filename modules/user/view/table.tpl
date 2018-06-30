<style>
    .table{
    margin-bottom: 10px;
    }
    .pagination{
    margin: 0px;
    }
</style>
<div class="row clearfix">
<div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
    <div class="card">
        <div class="header">
            <h2>
                <b>用户账号列表</b>
                <small>Add <code>.table-bordered</code> for borders on all sides of the table and cells.</small>
            </h2>
            <ul class="header-dropdown m-r--5">
                <li class="dropdown">
                    <a href="javascript:void(0);" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        <i class="material-icons">more_vert</i>
                    </a>
                    <ul class="dropdown-menu pull-right">
                        <li><a href="javascript:void(0);">Action</a></li>
                        <li><a href="javascript:void(0);">Another action</a></li>
                        <li><a href="javascript:void(0);">Something else here</a></li>
                    </ul>
                </li>
            </ul>
        </div>
        <div class="body table-responsive">
            <nav style="margin-bottom: 20px;">
                <button type="button" class="btn bg-teal waves-effect btn-xs" href="/user/add.html" onclick="AtContentOpen(this)">
                    <i class="material-icons">add_circle_outline</i>
                    <span>新的</span>
                </button>
                <button type="button" class="btn bg-red waves-effect btn-xs">
                    <i class="material-icons">clear</i>
                    <span>移除选中</span>
                </button>
                <button type="button" class="btn bg-red waves-effect btn-xs">
                    <i class="material-icons">block</i>
                    <span>禁用选中</span>
                </button>
                <button type="button" class="btn bg-red waves-effect btn-xs">
                    <i class="material-icons">check</i>
                    <span>启用选中</span>
                </button>
            </nav>
            <table class="table table-bordered">
                <thead>
                    <tr>
                        <th class="col-md-1" style="min-width: 120px">
                            <div class="unit">
                                <input type="checkbox" id="basic_checkbox_h" class="chk-col-red"/>
                                <label for="basic_checkbox_h">#</label>
                            </div>
                        </th>
                        <th>账号</th>
                        <th>角色</th>
                        <th>密码</th>
                        <th class="col-md-1" style="min-width: 168px">操作</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .list}}
                    <tr>
                        <th scope="row">
                            <div class="unit">
                                <input type="checkbox" id="basic_checkbox_{{.Id}}"  class="chk-col-red" />
                                <label for="basic_checkbox_{{.Id}}">{{.Id}}</label>
                            </div>
                        </th>
                        <td>{{.Name}}</td>
                        <td>{{.Role}}</td>
                        <td>{{.Passwd}}</td>
                        <td>
                            <div class="unit">
                                <label class="xbutton indigo">
                                    <i class="material-icons">insert_comment</i>
                                    <span>修改</span>
                                </label>
                                <span>/</span>
                                <label class="xbutton indigo">
                                    <i class="material-icons">block</i>
                                    <span>禁用</span>
                                </label>
                                <span>/</span>
                                <label class="xbutton red">
                                    <i class="material-icons">clear</i>
                                    <span>移除</span>
                                </label>
                            </div>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>

            <nav>
                <ul class="pagination pagination-sm">
                    <li>
                        <a href="javascript:void(0);" class="waves-effect">
                            <i class="material-icons">chevron_left</i>
                        </a>
                    </li>
                    {{pagination .page .n `<li><a href="/user/table?n={{.n}}" onclick="return AtContentOpen(this)" class="waves-effect">{{.n}}</a></li>`}}
                    <li>
                        <a href="javascript:void(0);" class="waves-effect">
                            <i class="material-icons">chevron_right</i>
                        </a>
                    </li>
                </ul>
            </nav>

        </div>
    </div>
</div>
</div>