<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Current Voter</h1>
     <p class="lead">Here you'll find all the available voters</p>
     {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Relation Name</th>
          <th>Gender</th>
          <th>District</th>
          <th>AC</th>
          <th>Section</th>
          <th>{{.counts}}</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Name_english}}</td>
          <td>{{$record.Relation_name_english}}</td>
          <td>{{$record.Gender}}</td>
          <td>{{$record.District_name_english}}</td>
          <td>{{$record.Ac_name_english}}</td>
          <td>{{$record.Section_name_english}}</td>
        </tr>
        {{end}}
      </tbody>
      <!--<tfoot>
        <tr>
          <td colspan="4"><a href="{{urlfor "ManageController.Add"}}" title="add new article">add new article</a></td>
        </tr>
      </tfoot>-->
    </table>
  </div>
</div>
