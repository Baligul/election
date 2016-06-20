<div class="container">
   <div class="row">
      <div class="hero-text">
         {{if .flash.error}}
         <blockquote>{{.flash.error}}</blockquote>
         {{end}}
         {{if .flash.notice}}
         <blockquote>{{.flash.notice}}</blockquote>
         {{end}}
      </div>
   </div>
   <div class="row">
      <div class="hero-text">
         <h2>Count of Voters</h2>
         <div class="row">
            <div class="col-xs-12">
               <div class="box box-info">
                  <!-- form start -->
                  <form class="form-horizontal">
                     <div class="box-body">
                        <div class="form-group">
                           <label for="voters" class="col-sm-2 control-label">Total Number of Voters: </label>
                           <div class="col-sm-10">
                              <label for="voters" class="col-sm-2 control-label">{{.VotersCount}}</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="muslimvoters" class="col-sm-2 control-label">Muslim Voters: </label>
                           <div class="col-sm-10">
                              <label for="muslimvoters" class="col-sm-2 control-label">{{.MuslimVotersCount}}</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="othervoters" class="col-sm-2 control-label">Other Voters: </label>
                           <div class="col-sm-10">
                              <label for="othervoters" class="col-sm-2 control-label">{{.OtherVotersCount}}</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="malevoters" class="col-sm-2 control-label">Male Voters: </label>
                           <div class="col-sm-10">
                              <label for="malevoters" class="col-sm-2 control-label">{{.MaleVotersCount}}</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="femalevoters" class="col-sm-2 control-label">Female Voters: </label>
                           <div class="col-sm-10">
                              <label for="femalevoters" class="col-sm-2 control-label">{{.FemaleVotersCount}}</label>
                           </div>
                        </div>
                     </div>
                     <!-- /.box-footer -->
                  </form>
               </div>
            </div>
            <!--/.col (right) -->
         </div>
      </div>
      <div class="hero-text">
         <h2>Percentage of Voters</h2>
         <div class="row">
            <div class="col-xs-12">
               <div class="box box-info">
                  <!-- form start -->
                  <form class="form-horizontal" method="post" action="/election/view">
                     <div class="box-body">
                        <div class="form-group">
                           <label for="voters" class="col-sm-2 control-label">Total Number of Voters: </label>
                           <div class="col-sm-10">
                              <label for="voters" class="col-sm-2 control-label">{{.VotersCount}}</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="muslimvoters" class="col-sm-2 control-label">Muslim Voters: </label>
                           <div class="col-sm-10">
                              <label for="muslimvoters" class="col-sm-2 control-label">{{.MuslimVotersPercentage}}%</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="othervoters" class="col-sm-2 control-label">Other Voters: </label>
                           <div class="col-sm-10">
                              <label for="othervoters" class="col-sm-2 control-label">{{.OtherVotersPercentage}}%</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="malevoters" class="col-sm-2 control-label">Male Voters: </label>
                           <div class="col-sm-10">
                              <label for="malevoters" class="col-sm-2 control-label">{{.MaleVotersPercentage}}%</label>
                           </div>
                        </div>
                        <div class="form-group">
                           <label for="femalevoters" class="col-sm-2 control-label">Female Voters: </label>
                           <div class="col-sm-10">
                              <label for="femalevoters" class="col-sm-2 control-label">{{.FemaleVotersPercentage}}%</label>
                           </div>
                        </div>
                     </div>
                     <!-- /.box-body -->
                     <div class="box-footer">
                        <a class="btn btn-default" href="/" style="padding-left: 11px;">Cancel</a>
                        <button type="submit" class="btn btn-danger pull-right">Continue</button>
                     </div>
                     <!-- /.box-footer -->
                  </form>
               </div>
            </div>
            <!--/.col (right) -->
         </div>
      </div>
   </div>
</div>