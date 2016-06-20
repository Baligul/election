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
      <div class="col-xs-12">
         <div class="col-md-12">
            <h1>Set base criteria for calculation of percentages</h1>
            <!-- form start -->
            <!--<p>{{if .status.Msg}}{{.status.Msg}}{{end}}</p> -->
            <form class="form-horizontal" method="post" action="/election/view">
               <input type="hidden" name="_csrf" value="{_csrf}" />
               <div class="form-group">
                  <label for="district"
                     class="col-md-1 control-label">District</label>
                  <div class="col-sm-10">
                     <select class="form-control select2"
                        id="district"
                        name="district"
                        multiple="multiple"
                        data-placeholder="Filter by District"
                        style="width: 100%;">
                        <option>Rampur</option>
                     </select>
                  </div>
               </div>
               <div class="form-group">
                  <label class="col-sm-2 control-label">Religion</label>
                  <div class="col-sm-10">
                     <div class="radio">
                        <label>
                        <input type="radio" name="religion" id="Muslim" value="Muslim"> Muslims
                        </label>
                     </div>
                     <div class="radio">
                        <label>
                        <input type="radio" name="religion" id="other" value="Other"> Others
                        </label>
                     </div>
                  </div>
               </div>
               <div class="form-group">
                  <label class="col-sm-2 control-label">Gender</label>
                  <div class="col-sm-10">
                     <div class="radio">
                        <label>
                        <input type="radio" name="gender" id="male" value="male"> Male
                        </label>
                     </div>
                     <div class="radio">
                        <label>
                        <input type="radio" name="gender" id="female" value="female"> Female
                        </label>
                     </div>
                  </div>
               </div>
               <div class="form-group">
                  <label for="acname"
                     class="col-md-1 control-label">AC / Kasba / City</label>
                  <div class="col-sm-10">
                     <select class="form-control select2"
                        id="acname"
                        name="acname"
                        multiple="multiple"
                        data-placeholder="Filter by AC Name"
                        style="width: 100%;">
                        <option>Rampur</option>
                        <option>Chamraua</option>
                        <option>Suar</option>
                        <option>Milak</option>
                        <option>Bilaspur</option>
                     </select>
                  </div>
               </div>
               <div class="form-group">
                  <label for="sectionname"
                     class="col-md-1 control-label">Section / Area / Mohalla</label>
                  <div class="col-sm-10">
                     <select class="form-control select2"
                        id="sectionname"
                        name="sectionname"
                        multiple="multiple"
                        data-placeholder="Filter by Section Name"
                        style="width: 100%;">
                        <option>[DHURYAI</option>
                        <option>001 CHAKFERI</option>
                        <option>2 AMEER ALI KHAN 1</option>
                        <option>AADPUR</option>
                        <option>AAGA-1</option>
                        <option>AAGA-2</option>
                        <option>AAKILPUR</option>
                        <option>AARSAL PARSAL-1</option>
                        <option>AARSAL PARSAL-2</option>
                        <option>AARSAL PARSAL-3</option>
                        <option>ABBAS NAGAR</option>
                        <option>ABBAS NAGAR SUAR</option>
                        <option>ABBASNAGAR TANDA</option>
                        <option>ABBASPUR</option>
                        <option>ABDULLANAGAR</option>
                        <option>ADILNAGAR</option>
                        <option>AFARIDIYAN-1</option>
                        <option>AFARIDIYAN-2</option>
                        <option>AFGANAN 01</option>
                        <option>AFGANAN 02</option>
                        <option>AFISARS BANGALE</option>
                        <option>AFZALPUR PATTI</option>
                        <option>AGAPUR</option>
                        <option>AGAPUR-1</option>
                        <option>AGAPUR-2</option>
                        <option>AGAPUR-3</option>
                        <option>AGAPUR-4</option>
                        <option>AGAPUR-5</option>
                        <option>AGAPUR-6</option>
                        <option>AGLAGA -1</option>
                        <option>AGLAGA -2</option>
                        <option>AHAMAD NAGAR PADMPUR</option>
                        <option>AHAMDABAD</option>
                        <option>AHAMDNAGAR NEAR ISMAIL NAGAR</option>
                        <option>AHAMDNAGAR NIKAT FAIZGANJ</option>
                        <option>AHMAD NAGAR JAGEER-1</option>
                        <option>AHMAD NAGAR JAGEER-2</option>
                        <option>AHMAD NAGAR KHEDA</option>
                        <option>AHMAD NAGAR NEAR BHAGWATIPUR</option>
                        <option>AHMAD NAGAR NIKAT JIWAI KADIM</option>
                        <option>AHMAD NAGAR PAHADI</option>
                        <option>AHMADABAD</option>
                        <option>AHMADABAD URF TELIPURA</option>
                        <option>AHMADGANJ URF MUKATPUR</option>
                        <option>AHMADNAGAR KA A</option>
                        <option>AHMADNAGAR NEAR AGA</option>
                        <option>AHMADNAGAR NEAR BHAGWANT NAGAR</option>
                        <option>AHMADNAGAR NIKAT KUNDANPUR</option>
                        <option>AHMADNAGAR TARANA</option>
                        <option>AHMADNAGAR THEGA</option>
                        <option>AHMADNAGAR URF MASAIYAN MOUHAMMDI</option>
                        <option>AHMADPUR</option>
                        <option>AHRAULA</option>
                        <option>AHRO-1</option>
                        <option>AHRO-2</option>
                        <option>AHRO-3</option>
                        <option>AHRO-4</option>
                        <option>AHROLA</option>
                        <option>AIMI-1</option>
                        <option>AIMI-2</option>
                        <option>AINCHORA-1</option>
                        <option>AINCHORA-2</option>
                        <option>AINJANKHEDA</option>
                        <option>AJAYPUR</option>
                        <option>AJEETPUR-1</option>
                        <option>AJEETPUR-10</option>
                        <option>AJEETPUR-2</option>
                        <option>AJEETPUR-3</option>
                        <option>AJEETPUR-4</option>
                        <option>AJEETPUR-5</option>
                        <option>AJEETPUR-6</option>
                        <option>AJEETPUR-7</option>
                        <option>AJEETPUR-8</option>
                        <option>AJITAPUR-9</option>
                        <option>AKAB MOTI MASJID</option>
                        <option>AKAB POST AFIS</option>
                        <option>AKHADA MALLI KHAN</option>
                        <option>AKHBARABAD</option>
                        <option>AKHUN KHE LAN</option>
                        <option>AKONDA</option>
                        <option>ALAFGUNJ</option>
                        <option>ALBANAGALA</option>
                        <option>ALI  NAGAR JAGEER-1</option>
                        <option>ALI  NAGAR JAGEER-2</option>
                        <option>ALI  NAGAR UTTARI</option>
                        <option>ALI NAGAR TAH</option>
                        <option>ALIGANJ BENJEER-1</option>
                        <option>ALIGANJ BENJEER-2</option>
                        <option>ALINAGAR JUNUBI</option>
                        <option>ALINAGAR KOTA</option>
                        <option>ALINAGAR SHUMALI</option>
                        <option>ALIPUR THEKA</option>
                        <option>ALLAPUR</option>
                        <option>AMBARPUR</option>
                        <option>AMEER ALI KHAN 02</option>
                        <option>ANDKHEDA</option>
                        <option>ANGUR VALI MASJID</option>
                        <option>ANGURI BAG</option>
                        <option>ANKHERA</option>
                        <option>ANWA 01</option>
                        <option>ANWA 02</option>
                        <option>ANWARIYA TALIWABAD</option>
                        <option>ANYA FAIKTARI ERIYA</option>
                        <option>ASALATPUR</option>
                        <option>ASALATPUR NEAR  SHIV NAGAR</option>
                        <option>ASHOKPUR</option>
                        <option>ASTABAL KAIMP MAY THANA GANJ</option>
                        <option>ATAI NAGAR</option>
                        <option>ATRIYA</option>
                        <option>ATTA ALLANUR</option>
                        <option>AVAS VIKAS GANGAPUR</option>
                        <option>AWADI RAILWAY STATION</option>
                        <option>AZEEM NAGAR</option>
                        <option>AZEEMNAGAR</option>
                        <option>BABOORA</option>
                        <option>BACHHIYAI</option>
                        <option>BADALI-1</option>
                        <option>BADALI-2</option>
                        <option>BADHPURA SHARKI</option>
                        <option>BADPURA</option>
                        <option>BADPURA SHUMALI</option>
                        <option>BAG PUKHTA</option>
                        <option>BAGADKHA-1</option>
                        <option>BAGADKHA-2</option>
                        <option>BAGADKHA-3</option>
                        <option>BAGI</option>
                        <option>BAGI-1</option>
                        <option>BAGI-2</option>
                        <option>BAGICHA AIMNA</option>
                        <option>BAGICHA CHHOTE MIAN 2</option>
                        <option>BAGICHA CHHOTE MIAN-1</option>
                        <option>BAGICHA GAJI MUJFAFAR KHA</option>
                        <option>BAGICHA JOKHIRAM</option>
                        <option>BAGRAWWA-1</option>
                        <option>BAGRAWWA-2</option>
                        <option>BAHADARGANJ</option>
                        <option>BAHADUR GANJ</option>
                        <option>BAHADURGANJ</option>
                        <option>BAHAPURA-2</option>
                        <option>BAHAPURI</option>
                        <option>BAHPUR GANGAPUR</option>
                        <option>BAINDU KHEDA</option>
                        <option>BAIRAM NAGAR</option>
                        <option>BAIZANI</option>
                        <option>BAJAR ABDULLA GANJ</option>
                        <option>BAJAR KALAN-1</option>
                        <option>BAJAR KALAN-2</option>
                        <option>BAJAR KALAN-3</option>
                        <option>BAJAR KALAN-4</option>
                        <option>BAJAR KALAN-5</option>
                        <option>BAJAR KALAN-6</option>
                        <option>BAJAR KALAN-7</option>
                        <option>BAJAR KALAN-8</option>
                        <option>BAJAR NASARULLA KHAN</option>
                        <option>BAJAR SAFDARGANJ MAY PAN DARIBA</option>
                        <option>BAJARIYA FATEH ALI KHA 2</option>
                        <option>BAJARIYA FATEH ALI KHA-1</option>
                        <option>BAJARIYA HIMMAT KHAN</option>
                        <option>BAJARIYA JHABBU KHAN MAY SILEYAN-1</option>
                        <option>BAJARIYA JHABBU KHAN MAY SILEYAN-2</option>
                        <option>BAJARIYA KADHU</option>
                        <option>BAJARIYA KHANASAMA</option>
                        <option>BAJARIYA MULLA JARIF</option>
                        <option>BAJHANPUR-01</option>
                        <option>BAJHANPUR-02</option>
                        <option>BAJODI TOLA -1</option>
                        <option>BAJODI TOLA -2</option>
                        <option>BAJODI TOLA- 4</option>
                        <option>BAJODI TOLA- 5</option>
                        <option>BAJODI TOLA-3</option>
                        <option>BAKENIYA</option>
                        <option>BAKENIYA BHAT</option>
                        <option>BAKNOARI</option>
                        <option>BALBHADRAPUR</option>
                        <option>BALKHEDA</option>
                        <option>BALUPURA</option>
                        <option>BAMANPURA-1</option>
                        <option>BAMANPURA-2</option>
                        <option>BAMANPURI</option>
                        <option>BAMNA</option>
                        <option>BAMNAPURI -1</option>
                        <option>BAMNAPURI -2</option>
                        <option>BANDAR 01</option>
                        <option>BANDAR 02</option>
                        <option>BANDUKACHIYAN</option>
                        <option>BANGALA AJAD KHA</option>
                        <option>BANGALA KASAM ALI KHAN</option>
                        <option>BANKRABAD</option>
                        <option>BANSKHERA</option>
                        <option>BANSNAGLI</option>
                        <option>BARA KHAS</option>
                        <option>BARA NIKAT GAJEJA</option>
                        <option>BARADARI</option>
                        <option>BARBALAN</option>
                        <option>BARDA GAUN 01</option>
                        <option>BARDA GAUN 02</option>
                        <option>BARDA GAUN 03</option>
                        <option>BARKHERA</option>
                        <option>BARUA -01</option>
                        <option>BARUA -02</option>
                        <option>BATHUAKHERA</option>
                        <option>BAZABALA</option>
                        <option>BAZABALA-1</option>
                        <option>BAZANA-2</option>
                        <option>BEDAN</option>
                        <option>BEDPUR</option>
                        <option>BEGMABAD</option>
                        <option>BEGMABAD NIKAT ASHOKPUR</option>
                        <option>BEGMABAD NIKAT KRIMCHA</option>
                        <option>BEHPURA- 1</option>
                        <option>BEHTA</option>
                        <option>BEHTRA</option>
                        <option>BELADARAN</option>
                        <option>BERIYAN AKAB MAHAL SARAY 1</option>
                        <option>BERIYAN AKAB MAHAL SARAY 2</option>
                        <option>BERKHEDA</option>
                        <option>BERKHEDI FEJABAD</option>
                        <option>BESRA-2</option>
                        <option>BHABANIPUR</option>
                        <option>BHAGBANTPUR</option>
                        <option>BHAGWANT NAGAR -2</option>
                        <option>BHAGWANT NAGAR-1</option>
                        <option>BHAGWATIPUR</option>
                        <option>BHAISAKHANA</option>
                        <option>BHAIYANGALA</option>
                        <option>BHAMROA-1</option>
                        <option>BHAMROA-2</option>
                        <option>BHAMROA-3</option>
                        <option>BHANDPURA-1</option>
                        <option>BHANDPURA-2</option>
                        <option>BHARATPUR</option>
                        <option>BHATIKHEDA</option>
                        <option>BHATPURA</option>
                        <option>BHATPURA CHAKRPAN</option>
                        <option>BHATPURA TARAN-1</option>
                        <option>BHATPURA TARAN-2</option>
                        <option>BHATTI TOLA-1</option>
                        <option>BHATTI TOLA-2</option>
                        <option>BHATTI TOLA-3</option>
                        <option>BHAUPURA</option>
                        <option>BHAWARKA-1</option>
                        <option>BHAWARKA-2</option>
                        <option>BHAWARKI</option>
                        <option>BHENSIYA JUALAPUR-1</option>
                        <option>BHENSIYA JUALAPUR-2</option>
                        <option>BHENSODI-3</option>
                        <option>BHENSORI-1</option>
                        <option>BHENSORI-2</option>
                        <option>BHITAR GAUN 01</option>
                        <option>BHITAR GAUN 02</option>
                        <option>BHOAJIPURA</option>
                        <option>BHOANAKPUR</option>
                        <option>BHOGPUR</option>
                        <option>BHOLANAGALA</option>
                        <option>BHOLAPUR KADEEM</option>
                        <option>BHOT</option>
                        <option>BHOT BAKKAL -1</option>
                        <option>BHOT BAKKAL -2</option>
                        <option>BHOT BAKKAL -3</option>
                        <option>BHOT BAKKAL -4</option>
                        <option>BHOT-2</option>
                        <option>BHOT-3</option>
                        <option>BHOUPATPUR</option>
                        <option>BHOURKA</option>
                        <option>BHOURKI JADID</option>
                        <option>BHOURKI KADEEM</option>
                        <option>BHUDASI 01</option>
                        <option>BHUDASI 02</option>
                        <option>BHUKSAURA</option>
                        <option>BHUVRA  MUSTAHAQAM-1</option>
                        <option>BHUVRA  MUSTAHAQAM-2</option>
                        <option>BHUVRA AHETMALI-1</option>
                        <option>BHUVRA AHETMALI-2</option>
                        <option>BIBRA</option>
                        <option>BICHPURI SANGRAMPUR</option>
                        <option>BIDAU</option>
                        <option>BIDPURA</option>
                        <option>BIDWA NAGALA</option>
                        <option>BIHARI NAGAR</option>
                        <option>BIHAT</option>
                        <option>BIJARKHATA UTTRI-1</option>
                        <option>BIJARKHATA UTTRI-2</option>
                        <option>BIJARKHATA UTTRI-3</option>
                        <option>BIJHDA</option>
                        <option>BIJLI</option>
                        <option>BIJPURI LALJI</option>
                        <option>BILASAPUR GATE</option>
                        <option>BISRA-1</option>
                        <option>BISRI</option>
                        <option>BIZPURI</option>
                        <option>BOODPUR</option>
                        <option>BOSENA</option>
                        <option>BRIJPUR</option>
                        <option>BUDANPUR</option>
                        <option>BUDHIYA KI DUKAN</option>
                        <option>BUDHPUR</option>
                        <option>BUDI DARIYAL</option>
                        <option>BULAND KVATARS</option>
                        <option>CHAH INCHHARAM</option>
                        <option>CHAH JADIYAN V JARGARAN</option>
                        <option>CHAH JATTA</option>
                        <option>CHAH KESRA SIH</option>
                        <option>CHAH KHAIYATAN</option>
                        <option>CHAH KHAJAN KHAN</option>
                        <option>CHAH KHANASAMA</option>
                        <option>CHAH MOTE KALLAN</option>
                        <option>CHAH SHOR-1</option>
                        <option>CHAH SHOR-2</option>
                        <option>CHAHASATAI</option>
                        <option>CHAHASHIRI</option>
                        <option>CHAINPUR</option>
                        <option>CHAK GAZROLA</option>
                        <option>CHAK HARDASPUR</option>
                        <option>CHAK KHARDIYA</option>
                        <option>CHAK KUNDESARI</option>
                        <option>CHAK LADPUR</option>
                        <option>CHAK NO-1 TANDA</option>
                        <option>CHAK NO-10(2) TANDA</option>
                        <option>CHAK NO-10(2)TANDA</option>
                        <option>CHAK NO-11 TANDA</option>
                        <option>CHAK NO-12 TANDA</option>
                        <option>CHAK NO-13 TANDA</option>
                        <option>CHAK NO-14 TANDA</option>
                        <option>CHAK NO-15 TANDA</option>
                        <option>CHAK NO-16 TANDA -1</option>
                        <option>CHAK NO-16TANDA-2</option>
                        <option>CHAK NO-17 TANDA</option>
                        <option>CHAK NO-18 (1)TANDA-1</option>
                        <option>CHAK NO-18 (2) TANDA</option>
                        <option>CHAK NO-18 TANDA</option>
                        <option>CHAK NO-19 TANDA</option>
                        <option>CHAK NO-2 TANDA</option>
                        <option>CHAK NO-20(1) TANDA</option>
                        <option>CHAK NO-20(2) TANDA</option>
                        <option>CHAK NO-21TANDA</option>
                        <option>CHAK NO-22(1)TANDA</option>
                        <option>CHAK NO-22(2) TANDA</option>
                        <option>CHAK NO-23 TANDA</option>
                        <option>CHAK NO-24TANDA</option>
                        <option>CHAK NO-25 (2)TANDA</option>
                        <option>CHAK NO-25 (3)TANDA</option>
                        <option>CHAK NO-25 TANDA</option>
                        <option>CHAK NO-25(1)-2TANDA</option>
                        <option>CHAK NO-25(3) TANDA-2</option>
                        <option>CHAK NO-26 TANDA-2</option>
                        <option>CHAK NO-26TANDA</option>
                        <option>CHAK NO-27TANDA</option>
                        <option>CHAK NO-28 TANDA</option>
                        <option>CHAK NO-3</option>
                        <option>CHAK NO-4 TANDA</option>
                        <option>CHAK NO-5 TANDA</option>
                        <option>CHAK NO-6 TANDA</option>
                        <option>CHAK NO-7 TANDA</option>
                        <option>CHAK NO-8 TANDA</option>
                        <option>CHAK NO-9- TANDA(1)</option>
                        <option>CHAK NO-9(2) TANDA</option>
                        <option>CHAK SUAR -1</option>
                        <option>CHAK SUAR -2</option>
                        <option>CHAK SUAR -3</option>
                        <option>CHAK SUAR -4</option>
                        <option>CHAK SUAR-5</option>
                        <option>CHAK SUAR-6</option>
                        <option>CHAK SUAR-7</option>
                        <option>CHAKARPUR</option>
                        <option>CHAKARPUR BHUD</option>
                        <option>CHAKARPUR KADEEM</option>
                        <option>CHAKARPUR KADEEM 02</option>
                        <option>CHAKASHADINGAR</option>
                        <option>CHAKASHI</option>
                        <option>CHAKFERI</option>
                        <option>CHAKIYA HAYATNAGAR-1</option>
                        <option>CHAKIYA HAYATNAGAR-2</option>
                        <option>CHAKJEEM</option>
                        <option>CHAMARAN -2</option>
                        <option>CHAMARAN-1</option>
                        <option>CHAMARPURA</option>
                        <option>CHAMRA</option>
                        <option>CHAMROL</option>
                        <option>CHAMROUA-1</option>
                        <option>CHAMROUA-2</option>
                        <option>CHAMROUA-3</option>
                        <option>CHAMROUA-4</option>
                        <option>CHAMROUA-5</option>
                        <option>CHAMROUA-6</option>
                        <option>CHANDAYAN</option>
                        <option>CHANDELA</option>
                        <option>CHANDPUR</option>
                        <option>CHANDPURA JADID</option>
                        <option>CHANDPURA KADIM</option>
                        <option>CHANDPURA KALAN -02</option>
                        <option>CHANDPURA KALAN-1</option>
                        <option>CHANDPURA KHURD</option>
                        <option>CHANDPURA SALIS</option>
                        <option>CHANDUPURA SHIKAMPUR-1</option>
                        <option>CHANDUPURA SHIKAMPUR-2</option>
                        <option>CHANDUPURA SHIKAMPUR-3</option>
                        <option>CHANDUPURI</option>
                        <option>CHANDWA NAGALA</option>
                        <option>CHAPARRA</option>
                        <option>CHARAKH VALI MASJID</option>
                        <option>CHATARPUR</option>
                        <option>CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-1</option>
                        <option>CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-2</option>
                        <option>CHAUKONI</option>
                        <option>CHAUPURA-1</option>
                        <option>CHAUPURA-2</option>
                        <option>CHAUPURA-3</option>
                        <option>CHAUPURA-4</option>
                        <option>CHHATRAPUR</option>
                        <option>CHHIPIYAN MAY BADHEYAN -1</option>
                        <option>CHHIPIYAN MAY BADHEYAN -2</option>
                        <option>CHHITARIYA GAJIR-1</option>
                        <option>CHHITARIYA GAJIR-2</option>
                        <option>CHHITARIYA GAJIR-3</option>
                        <option>CHHITARIYA GAJIR-4</option>
                        <option>CHICHOLI</option>
                        <option>CHIKNA</option>
                        <option>CHIKTI RAMNAGAR-1</option>
                        <option>CHIKTI RAMNAGAR-2</option>
                        <option>CHIRAN</option>
                        <option>CHITOUNI-1</option>
                        <option>CHITOUNI-2</option>
                        <option>CHOUKONI</option>
                        <option>CHUKHANDI</option>
                        <option>DABKA-1</option>
                        <option>DABKA-2</option>
                        <option>DADHIYAL EHATMALI-1</option>
                        <option>DADHIYAL EHATMALI-2</option>
                        <option>DADHIYAL EHATMALI-3</option>
                        <option>DADHIYAL EHATMALI-4</option>
                        <option>DADHIYAL EHATMALI-5</option>
                        <option>DADHIYAL EHATMALI-6</option>
                        <option>DADHIYAL MUSTEKAM -1</option>
                        <option>DADHIYAL MUSTEKAM -2</option>
                        <option>DADHIYAL MUSTEKAM -3</option>
                        <option>DADHIYAL MUSTEKAM-4</option>
                        <option>DADHIYAL MUSTEKAM-5</option>
                        <option>DADHIYAL MUSTEKAM-6</option>
                        <option>DADHIYAL MUSTEKAM-7</option>
                        <option>DAKKANAGALIYA</option>
                        <option>DAKURIYA</option>
                        <option>DALELNAGAR-1</option>
                        <option>DALELNAGAR-2</option>
                        <option>DALKI</option>
                        <option>DANDIYA NAYAMATGANG</option>
                        <option>DANIAPUR</option>
                        <option>DANKARA</option>
                        <option>DANKARI</option>
                        <option>DANYAPUR SHANKARPUR</option>
                        <option>DARA NAGAR</option>
                        <option>DARAKHT KAITH</option>
                        <option>DARAKHT KHINNI</option>
                        <option>DARAV NAGAR</option>
                        <option>DARIYA GADH</option>
                        <option>DARSHANPUR</option>
                        <option>DAUKPURI TANDA -1</option>
                        <option>DAUKPURI TANDA-2</option>
                        <option>DAUKPURI TANDA-3</option>
                        <option>DAUKPURI TANDA-4</option>
                        <option>DEENPUR</option>
                        <option>DEHALI DARAVAJA MAY GAUKHANA KAIMP</option>
                        <option>DEVIPURA</option>
                        <option>DEVRANIYA SHUMALI 1</option>
                        <option>DEVRANIYA SHUMALI 2</option>
                        <option>DEVRI BUJURG</option>
                        <option>DEVRI KHURD</option>
                        <option>DEWRANIYA SHARKI</option>
                        <option>DHAKI</option>
                        <option>DHAKIYA 02</option>
                        <option>DHAKIYA 03</option>
                        <option>DHAKIYA 04</option>
                        <option>DHAKIYA 1</option>
                        <option>DHAKKA HAJINAGAR 1</option>
                        <option>DHAKKA HAJINAGAR-2</option>
                        <option>DHAKPURA</option>
                        <option>DHALLIYA</option>
                        <option>DHAMORA-1</option>
                        <option>DHAMORA-2</option>
                        <option>DHAMORA-3</option>
                        <option>DHANELI POORVI-1</option>
                        <option>DHANELI POORVI-2</option>
                        <option>DHANELI UTTARI-1</option>
                        <option>DHANELI UTTARI-2</option>
                        <option>DHANOORA</option>
                        <option>DHANOORI-1</option>
                        <option>DHANORA</option>
                        <option>DHANOURA</option>
                        <option>DHANOURI-2</option>
                        <option>DHANPUR NIKAT  SHAHDRA</option>
                        <option>DHANUPURA</option>
                        <option>DHANUPURA-1</option>
                        <option>DHANUPURA-2</option>
                        <option>DHANUPURA-3</option>
                        <option>DHARAMPUR</option>
                        <option>DHARAMPUR UTTARI</option>
                        <option>DHARAMPURA</option>
                        <option>DHARMASHALA MAY MAI KA THAN</option>
                        <option>DHARMPUR</option>
                        <option>DHARMPUR GARVI</option>
                        <option>DHAWANI BUJURG-1</option>
                        <option>DHAWANI BUJURG-2</option>
                        <option>DHAWNI HASANPUR-1</option>
                        <option>DHAWNI HASANPUR-2</option>
                        <option>DHIMARKHEDA</option>
                        <option>DHIMRI</option>
                        <option>DHIRAZNAGAR</option>
                        <option>DHNORA</option>
                        <option>DHOLSAR</option>
                        <option>DHURIYAI</option>
                        <option>DIBDIBA-1</option>
                        <option>DIBDIBA-2</option>
                        <option>DIBDIBA-3</option>
                        <option>DIBDIBA-4</option>
                        <option>DIBDIBA-5</option>
                        <option>DILARI</option>
                        <option>DILPURA-1</option>
                        <option>DILPURA-2</option>
                        <option>DIVAAN</option>
                        <option>DIVIYANAGLA</option>
                        <option>DOHAIYA</option>
                        <option>DOHARIYA</option>
                        <option>DOMAHALA ROD MAY GHER GULAM NASIR KHAN</option>
                        <option>DUABAT</option>
                        <option>DUDAI-2</option>
                        <option>DUGANPUR</option>
                        <option>DUKAN SONA BINNA</option>
                        <option>DULICHANDPUR</option>
                        <option>DUNDAI-1</option>
                        <option>DUNDAWALA  AEHATMALI , MUSTEHAQAM-1</option>
                        <option>DUNDAWALA  AEHATMALI , MUSTEHAQAM-2</option>
                        <option>DUNDAWALA AEHATMALI , MUSTEHKAM 3</option>
                        <option>DUNDAWALA AEHATMALI , MUSTEHKAM 4</option>
                        <option>DUNDAWALA AEHATMALI , MUSTEHKAM 5</option>
                        <option>DUNGARPUR</option>
                        <option>DURJANPUR</option>
                        <option>DURNAGLA</option>
                        <option>FADASHEKH MAY BANS MANDI</option>
                        <option>FAIJGANJ</option>
                        <option>FAIJNAGAR</option>
                        <option>FAIZNAGAR</option>
                        <option>FAIZULLANAGAR-1</option>
                        <option>FAIZULLANAGAR-2</option>
                        <option>FAJALPUR</option>
                        <option>FAJALPUR NEAR DHANOORY</option>
                        <option>FAJULANAGAR</option>
                        <option>FAKEERGANJ URF GHOSIPURA</option>
                        <option>FARRAASHAN 01</option>
                        <option>FARRAASHAN 02</option>
                        <option>FARRAASHAN 03</option>
                        <option>FARRAASHAN 04</option>
                        <option>FATTEPUR</option>
                        <option>FAZALPUR PUSWADA</option>
                        <option>FAZGANJ</option>
                        <option>FILAKHANA JADID</option>
                        <option>FOOLPUR</option>
                        <option>FULAVAD</option>
                        <option>FULSUNGHA</option>
                        <option>FULSUNGHI</option>
                        <option>FUTAMAHAL MAY JAIN MANDIR</option>
                        <option>GADA</option>
                        <option>GADAIYA NIKAT HARSONAGLA</option>
                        <option>GADAIYA NIKAT HINGANAGLA</option>
                        <option>GADDI NAGLI- 2</option>
                        <option>GADDINAGLI-1</option>
                        <option>GADIYA NASEEMGANJ</option>
                        <option>GADMAR PATTI TIKA SINGH</option>
                        <option>GADMARPATTI TIKA SINGH</option>
                        <option>GAGAN NAGALA</option>
                        <option>GAGNPUR</option>
                        <option>GAJEJJA</option>
                        <option>GAJRAULA</option>
                        <option>GAJUPURA</option>
                        <option>GALPURA</option>
                        <option>GAMMANAPURA</option>
                        <option>GANGAN NAGLI</option>
                        <option>GANGAPUR JADID</option>
                        <option>GANGAPUR KADIM-1</option>
                        <option>GANGAPUR KADIM-2</option>
                        <option>GANGAPUR KALN</option>
                        <option>GANGAPURSHARKI</option>
                        <option>GANJ KADIM -2</option>
                        <option>GANJAKADIM -1</option>
                        <option>GAUKHANA KADIM</option>
                        <option>GAZROLA</option>
                        <option>GEHLUIYA</option>
                        <option>GEHNI</option>
                        <option>GHAIR AJAM KHAN</option>
                        <option>GHAIR ALLI KHA</option>
                        <option>GHAIR BAKHSHI</option>
                        <option>GHAIR DARIYA KHA</option>
                        <option>GHAIR GILJIYAN</option>
                        <option>GHAIR KUTVI MIYAN</option>
                        <option>GHAIR MILKIYAN</option>
                        <option>GHAIR MULLA MALUK</option>
                        <option>GHAIR MUNSHIVALA MAY  CHAH SOTIYAN</option>
                        <option>GHAIR SAIFAL KHAN</option>
                        <option>GHAIR SHAH MOHAMMAD KHAN</option>
                        <option>GHAIR SHAIRFUDDIN KHAN</option>
                        <option>GHANASHYAMAPUR</option>
                        <option>GHARAMPUR</option>
                        <option>GHER AJIM KHA MAY AKHUNAJADA</option>
                        <option>GHER ALAF KHAN</option>
                        <option>GHER BAJ KHA MAY FILAVALAN-1</option>
                        <option>GHER BAJ KHA MAY FILAVALAN-2</option>
                        <option>GHER BAJ KHA MAY FILAVALAN-3</option>
                        <option>GHER BECHA KHAN</option>
                        <option>GHER CHANDAN KHAN MAY JEL K VATARS</option>
                        <option>GHER FARUKHSHAH KHAN</option>
                        <option>GHER FATEH KHAN</option>
                        <option>GHER GAUS MOHAMMAD KHA</option>
                        <option>GHER HASSAN KHA</option>
                        <option>GHER INAYAT KHAN</option>
                        <option>GHER JANAS KHAN</option>
                        <option>GHER JIYAUNNAVI KHA</option>
                        <option>GHER KALANDAR KHAN-1</option>
                        <option>GHER KALANDAR KHAN-2</option>
                        <option>GHER KATE BAJ KHAN</option>
                        <option>GHER MARDAN KHAN</option>
                        <option>GHER MEERBAZ KHAN MEY GHER HASSN KHAN-1</option>
                        <option>GHER MEERBAZ KHAN MEY GHER HASSN KHAN-2</option>
                        <option>GHER MIAN KHAN -2</option>
                        <option>GHER MIAN KHAN -3</option>
                        <option>GHER MIAN KHAN-1</option>
                        <option>GHER MIYAN KHEL</option>
                        <option>GHER MOHABBAT KHA</option>
                        <option>GHER MUBARAK SHAH KHAN</option>
                        <option>GHER MUHAB KHA</option>
                        <option>GHER MULLA GAINRAT</option>
                        <option>GHER MUNAVVAR KHAN</option>
                        <option>GHER NAJJU KHAN</option>
                        <option>GHER PIPALWALA-1</option>
                        <option>GHER PIPALWALA-2</option>
                        <option>GHER PURAN SINGH</option>
                        <option>GHER REHMAT KHAN 1</option>
                        <option>GHER REHMAT KHAN 2</option>
                        <option>GHER SAIFUDADIN KHAN</option>
                        <option>GHER SAIFUDADIN KHAN 2</option>
                        <option>GHER SALAVAT KHAN</option>
                        <option>GHER SARABDAL KHAN</option>
                        <option>GHER TONGA-1</option>
                        <option>GHER TONGA-2</option>
                        <option>GHER TONGA-3</option>
                        <option>GHER USMAN KHA</option>
                        <option>GHER YUSUF KHAN</option>
                        <option>GHOR SAKHI</option>
                        <option>GIRDHARPUR</option>
                        <option>GODHI-1</option>
                        <option>GODHI-2</option>
                        <option>GOHARPUR</option>
                        <option>GOKAL NAGRI</option>
                        <option>GOSIPUR</option>
                        <option>GOUSAMPUR -2</option>
                        <option>GOUSAMPUR 01</option>
                        <option>GOVIND PURA NEAR PADAMPUR</option>
                        <option>GOVINDPURA</option>
                        <option>GOVINDPURA NEAR RAYPUR</option>
                        <option>GUJAR TOLA</option>
                        <option>GUJRAILA</option>
                        <option>GULAD PEEPALSANA-1</option>
                        <option>GULAD PEEPALSANA-2</option>
                        <option>GULADIYA BHAT</option>
                        <option>GULADIYA BIJLA</option>
                        <option>GULADIYA TEULA</option>
                        <option>GULALPUR</option>
                        <option>GULAMGANJ</option>
                        <option>GULARIYA KALAN</option>
                        <option>HAFIJNAGAR</option>
                        <option>HAJINAGAR</option>
                        <option>HAJRATPUR</option>
                        <option>HAKIMAAN 01</option>
                        <option>HAKIMAAN 02</option>
                        <option>HAKIMGANJ</option>
                        <option>HALUNAAGAR</option>
                        <option>HAMEERPUR</option>
                        <option>HAMIDABAD-1</option>
                        <option>HAMIDABAD-2</option>
                        <option>HAMIDNAGAR</option>
                        <option>HAMIR PUR -3</option>
                        <option>HAMIRPUR-1</option>
                        <option>HAMIRPUR-2</option>
                        <option>HAMIRPUR-4</option>
                        <option>HAMIRPUR-5</option>
                        <option>HAMIRPUR-6</option>
                        <option>HARAIYA KURD</option>
                        <option>HARDASPUR</option>
                        <option>HARDASPUR KOTRA-1</option>
                        <option>HARDASPUR KOTRA-2</option>
                        <option>HARDUA</option>
                        <option>HARETA-1</option>
                        <option>HARETA-2</option>
                        <option>HAREYYA</option>
                        <option>HAREYYA KA MAJHRA</option>
                        <option>HARIYA KALAN</option>
                        <option>HARIYALE</option>
                        <option>HARJIPUR</option>
                        <option>HARNAGLA</option>
                        <option>HARSONAGLA</option>
                        <option>HASAN PUR UTTARI</option>
                        <option>HASANPUR BILASPUR</option>
                        <option>HASANPUR SHARKI</option>
                        <option>HASHMINAGAR</option>
                        <option>HASNAPUR</option>
                        <option>HATA KALAN-3</option>
                        <option>HAYATNAGAR DAKSHNI</option>
                        <option>HIMMATGANJ</option>
                        <option>HIMMATPUR 01</option>
                        <option>HIMMATPUR 02</option>
                        <option>HINGANAGLA</option>
                        <option>HIRANKHERA</option>
                        <option>HOSPUR</option>
                        <option>HURMATNAGAR</option>
                        <option>HURMATNAGAR TANDA -1</option>
                        <option>HURMATNAGAR TANDA -2</option>
                        <option>HUSAIN GANJ BHUDA</option>
                        <option>HUSAINGANJ</option>
                        <option>HUSAINGANJ GADDINAGLI</option>
                        <option>HUSAINGANJ NEAR JANAKPUR</option>
                        <option>HUSHIANGUNJ NEAR MADARPUR</option>
                        <option>HYATNAGAR</option>
                        <option>IMAMBADA -1</option>
                        <option>IMAMBADA -2</option>
                        <option>IMAMBADA -4</option>
                        <option>IMAMBADA -5</option>
                        <option>IMAMBADA-3</option>
                        <option>IMARATI NIKAT AKBRABAD</option>
                        <option>IMARTI</option>
                        <option>IMLI ASMAT KHAN</option>
                        <option>IMLI BATANIYA</option>
                        <option>IMLI JHULEWALI</option>
                        <option>IMRATA RAY</option>
                        <option>IMRATA-1</option>
                        <option>IMRATA-2</option>
                        <option>IMRATA-3</option>
                        <option>IMRATA-4</option>
                        <option>IMRATA-5</option>
                        <option>IMRATPUR</option>
                        <option>INAYATPUR</option>
                        <option>INDARPUR</option>
                        <option>INDRA-1</option>
                        <option>INDRI</option>
                        <option>ISAKHERA</option>
                        <option>ISAMAIL NAGAR NEAR AHMADNAGAR</option>
                        <option>ISWARPUR</option>
                        <option>ITANGA BERAMNAGAR</option>
                        <option>JADOANPUR-2</option>
                        <option>JADONPUR-1</option>
                        <option>JAFRABAD</option>
                        <option>JAGANNATHPUR</option>
                        <option>JAGATPUR</option>
                        <option>JAGDISHPUR</option>
                        <option>JAGESAR</option>
                        <option>JAHANGEERPUR</option>
                        <option>JAIDOLI</option>
                        <option>JAITOLI 01</option>
                        <option>JAITOLI 02</option>
                        <option>JALAFNAGLA-1</option>
                        <option>JALAFNAGLA-2</option>
                        <option>JALIF NAGLA-1</option>
                        <option>JALIF NAGLA-2</option>
                        <option>JALIF NAGLA-3</option>
                        <option>JAMANPUR</option>
                        <option>JAMAPUR</option>
                        <option>JAMIN GANJ</option>
                        <option>JANAKPUR</option>
                        <option>JANUNAGAR</option>
                        <option>JASMOLI-1</option>
                        <option>JASMOLI-2</option>
                        <option>JATPURA</option>
                        <option>JATPURA-1</option>
                        <option>JATPURA-2</option>
                        <option>JAULAPUR</option>
                        <option>JAURASI</option>
                        <option>JHANDA BADE PIR SAHAB-1</option>
                        <option>JHANDA BADE PIR SAHAB-2</option>
                        <option>JHUNAIYA</option>
                        <option>JHURAKJHUNDI</option>
                        <option>JILEDARAN 01</option>
                        <option>JILEDARAN 02</option>
                        <option>JINA INAYAT KHA</option>
                        <option>JITHANIYA JAGIR</option>
                        <option>JITHANYA SHARKI</option>
                        <option>JIWAI JADID</option>
                        <option>JIWAI KADIM-1</option>
                        <option>JIWAI KADIM-2</option>
                        <option>JIYARAT SHIRIMIYAN-1</option>
                        <option>JIYARAT SHIRIMIYAN-2</option>
                        <option>JIYARAT SHIRIMIYAN-3</option>
                        <option>JIYARAT SHIRIMIYAN-4</option>
                        <option>JIYARAT SHIRIMIYAN-5</option>
                        <option>JIYARAT SHIRIMIYAN-6</option>
                        <option>JMALPUR</option>
                        <option>JUTHIYA-1</option>
                        <option>JUTHIYA-2</option>
                        <option>JUTHIYA-3</option>
                        <option>JVALANGAR AI0 D0 OII0 0 LE0KA0-5</option>
                        <option>JVALANGAR C 0 OII0 D0 V LEBAR KALONI-3</option>
                        <option>JVALANGAR C 0 OII0 D0 V LEBAR KALONI-4</option>
                        <option>JVALANGAR C0 OII0 D 0 LEVAR KALONI-1</option>
                        <option>JVALANGAR C0 OII0 D0 LEVAR KALONI-6</option>
                        <option>JVALANGAR C0 OII0 D0 LEVAR KALONI-7</option>
                        <option>JVALANGAR C0 OII0 D0 LEVAR KALONI-9</option>
                        <option>JVALANGAR C0 OII0 D0 V LEBAR KALONI-2</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-10</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-11</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-12</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-13</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-14</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-15</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-16</option>
                        <option>JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-17</option>
                        <option>JVALANGAR C0OII0D0 V LEVAR KALONI -8</option>
                        <option>JYARAT HALKE VALI</option>
                        <option>JYOHRA-1</option>
                        <option>JYOHRA-2</option>
                        <option>KACHNAL</option>
                        <option>KADRIGANJ</option>
                        <option>KAGANAGLA</option>
                        <option>KAISONAGALI TANDA</option>
                        <option>KAITHOLA</option>
                        <option>KAJARHAI</option>
                        <option>KAJIYAPURA</option>
                        <option>KAKKROA-3</option>
                        <option>KAKROA-1</option>
                        <option>KAKROA-2</option>
                        <option>KAKROA-4</option>
                        <option>KAKROA-5</option>
                        <option>KALAYANPUR</option>
                        <option>KALGHAR-1</option>
                        <option>KALGHAR-2</option>
                        <option>KALIYA NAGALA</option>
                        <option>KALKATTA</option>
                        <option>KALRAKH</option>
                        <option>KALYANPUR</option>
                        <option>KALYANPUR-1</option>
                        <option>KALYANPUR-2</option>
                        <option>KAMALPUR</option>
                        <option>KAMORA-1</option>
                        <option>KAMORA-2</option>
                        <option>KAMRUDDIN NAGAR</option>
                        <option>KAMUA NAGALA</option>
                        <option>KANAKPUR-1</option>
                        <option>KANAKPUR-2</option>
                        <option>KANOONGOYAN 01</option>
                        <option>KANOONGOYAN 02</option>
                        <option>KANOONGOYAN 03</option>
                        <option>KAPNERI</option>
                        <option>KARAITHI -1</option>
                        <option>KARAITHI-2</option>
                        <option>KARANPUR</option>
                        <option>KAREEMGUNJ</option>
                        <option>KAREEMPUR</option>
                        <option>KAREEMPUR SHARKI</option>
                        <option>KARIMGANJ</option>
                        <option>KARIMPUR GARVI</option>
                        <option>KARINGA</option>
                        <option>KARKHEDA-1</option>
                        <option>KARKHEDA-2</option>
                        <option>KARKHEDI</option>
                        <option>KARSAULA</option>
                        <option>KARSAULI</option>
                        <option>KASAMGANJ</option>
                        <option>KASHAVNAGLA</option>
                        <option>KASHIPUR -1</option>
                        <option>KASHIPUR -2</option>
                        <option>KASHIPUR-1</option>
                        <option>KASHIPUR-2</option>
                        <option>KASHIPUR-3</option>
                        <option>KASHIPUR-4</option>
                        <option>KASHIPUR-5</option>
                        <option>KASHIPUR-6</option>
                        <option>KASHIPUR-7</option>
                        <option>KASHIPUR-8</option>
                        <option>KASIYA KUNDA-1</option>
                        <option>KASIYA KUNDA-2</option>
                        <option>KASSHABAN</option>
                        <option>KASWA RAJPUR</option>
                        <option>KATAKUIYA MAY FILAKHANA KADIM-1</option>
                        <option>KATAKUIYA MAY FILAKHANA KADIM-2</option>
                        <option>KATRA JALALUDIN MAY GHER HASSAN KHAN</option>
                        <option>KAUSHALGANJ</option>
                        <option>KAUTALKHEDA</option>
                        <option>KAYAMGANJ</option>
                        <option>KAYASTHAN-1</option>
                        <option>KAYASTHAN-2</option>
                        <option>KEORAR-1</option>
                        <option>KEORAR-2</option>
                        <option>KESARPUR</option>
                        <option>KESARPUR NEAR KAREEMGANJ</option>
                        <option>KESHARPUR NEAR RAYPUR</option>
                        <option>KESHONGLI SUAR</option>
                        <option>KEWALPUR</option>
                        <option>KHAIRULLA PUR</option>
                        <option>KHAJURIYA</option>
                        <option>KHAJURIYA KALAN</option>
                        <option>KHAJURIYA KHURD</option>
                        <option>KHAMARIYA-1</option>
                        <option>KHAMARIYA-2</option>
                        <option>KHAMRI</option>
                        <option>KHANDELI</option>
                        <option>KHANDELI 02</option>
                        <option>KHANDI KHEDA</option>
                        <option>KHANDIA-1</option>
                        <option>KHANDIA-2</option>
                        <option>KHANDIA-3</option>
                        <option>KHANDIYA</option>
                        <option>KHANJIPURA</option>
                        <option>KHANPUR</option>
                        <option>KHANPUR JADID-1</option>
                        <option>KHANPUR JADID-2</option>
                        <option>KHANPUR KADIM</option>
                        <option>KHANPUR KHARVI</option>
                        <option>KHANPUR UTTARI-1</option>
                        <option>KHANPUR UTTARI-2</option>
                        <option>KHANUPURA</option>
                        <option>KHARADIYAN</option>
                        <option>KHARDIYA</option>
                        <option>KHARSOUL 01</option>
                        <option>KHARSOUL 02</option>
                        <option>KHATA</option>
                        <option>KHATA CHINTAMAN</option>
                        <option>KHATA KALAN-1</option>
                        <option>KHATA KALAN-2</option>
                        <option>KHATA KALAN-4</option>
                        <option>KHATKAN PASIYAN V TALAV MASTU KHAN</option>
                        <option>KHATPUR</option>
                        <option>KHAUD-1</option>
                        <option>KHAUD-2</option>
                        <option>KHAUNDALPUR-1</option>
                        <option>KHAUNDALPUR-2</option>
                        <option>KHAZUYA KHEDA</option>
                        <option>KHEDA GAJROLA</option>
                        <option>KHEDATANDA-1</option>
                        <option>KHEDATANDA-2</option>
                        <option>KHEDATANDA-3</option>
                        <option>KHEDATANDA-4</option>
                        <option>KHEMPUR</option>
                        <option>KHEMPUR-1</option>
                        <option>KHEMPUR-2</option>
                        <option>KHERA</option>
                        <option>KHERA PARSAL</option>
                        <option>KHIMOTIYA BAKHTI</option>
                        <option>KHIMOTIYA KHEDA</option>
                        <option>KHIRKA</option>
                        <option>KHIZARPUR</option>
                        <option>KHOD KALA-1</option>
                        <option>KHOD KALA-2</option>
                        <option>KHODPURA 1</option>
                        <option>KHODPURA 2</option>
                        <option>KHSNDSAR KOHNA</option>
                        <option>KHUBIYA NAGALA</option>
                        <option>KHUDAGANJ</option>
                        <option>KHUDANAGAR</option>
                        <option>KHUNTA KHERA</option>
                        <option>KHURSHEED NAGAR</option>
                        <option>KHUSHHALPUR</option>
                        <option>KHUSRO BAG</option>
                        <option>KHUTIA</option>
                        <option>KILA FORT</option>
                        <option>KIRA 01</option>
                        <option>KIRA 02</option>
                        <option>KIRA 03</option>
                        <option>KISANPUR ATRIYA</option>
                        <option>KISANPUR PANCHAKKI</option>
                        <option>KISHANPUR DULLANAGALA</option>
                        <option>KISHANPUR MOLLAGARH</option>
                        <option>KISHNPUR</option>
                        <option>KISROL</option>
                        <option>KITAPLAI</option>
                        <option>KOOP -2</option>
                        <option>KOOP -3</option>
                        <option>KOOP-1</option>
                        <option>KOTA ALINAGAR</option>
                        <option>KOTAVALAN</option>
                        <option>KOTAVALI KOHANA</option>
                        <option>KOTHA JAGEER</option>
                        <option>KOTHI KHAS BAG-1</option>
                        <option>KOTHI KHAS BAG-2</option>
                        <option>KOTRA</option>
                        <option>KOYLA-1</option>
                        <option>KOYLA-2</option>
                        <option>KOYLA-3</option>
                        <option>KOYLA-4</option>
                        <option>KOYLI</option>
                        <option>KRIMCHA-1</option>
                        <option>KRIMCHA-2</option>
                        <option>KRIMCHA-3</option>
                        <option>KRIPYA HAPPU</option>
                        <option>KRIPYA PANDE</option>
                        <option>KUAKHERA</option>
                        <option>KUCHA DEVIDAS</option>
                        <option>KUCHA PARAMESHVARI DAS</option>
                        <option>KUCHITA</option>
                        <option>KUIYA</option>
                        <option>KUKDI KHEDA</option>
                        <option>KUMHARIYA KHURD</option>
                        <option>KUMHARIYA KLA</option>
                        <option>KUMHARIYA-1</option>
                        <option>KUMHARIYA-2</option>
                        <option>KUNCHA ATISH BAJAN</option>
                        <option>KUNCHA BHAGMAL</option>
                        <option>KUNCHA FIRANGAN</option>
                        <option>KUNCHA KAJI</option>
                        <option>KUNCHA KASSAVAKHANA</option>
                        <option>KUNCHA LALA MIYAN</option>
                        <option>KUNCHA LANGARAKHANA</option>
                        <option>KUNCHA NAKKALAN</option>
                        <option>KUNCHA NALVANDAN</option>
                        <option>KUNCHA SEEKALGARANN</option>
                        <option>KUNDA 1</option>
                        <option>KUNDA 3</option>
                        <option>KUNDESRA</option>
                        <option>KUNDESRI</option>
                        <option>KUNVARAPUR NANAKAR-1</option>
                        <option>KUNVARAPUR NANAKAR-2</option>
                        <option>KUNVARPUR PADAMAPUR</option>
                        <option>KURTHIYA</option>
                        <option>LADORA NARAYANPUR</option>
                        <option>LADORI</option>
                        <option>LADPUR</option>
                        <option>LADPUR BIB</option>
                        <option>LADPUR NEAR BATHUAKHERA</option>
                        <option>LADPUR NEAR SEMRA-1</option>
                        <option>LADPUR NEAR SEMRA-2</option>
                        <option>LAKHIMPUR BHIKA</option>
                        <option>LAKHIMPUR VISHNU</option>
                        <option>LAKHMAN NAGALA</option>
                        <option>LAKHMIPUR</option>
                        <option>LAKHNAKHEDA</option>
                        <option>LAL KABAR-1</option>
                        <option>LAL KABAR-2</option>
                        <option>LAL MASJID-1</option>
                        <option>LAL MASJID-2</option>
                        <option>LALA NAGLA</option>
                        <option>LALPUR</option>
                        <option>LALPUR KALA -1</option>
                        <option>LALPUR KALA -2</option>
                        <option>LALPUR KALA -3</option>
                        <option>LALPUR KALA-4</option>
                        <option>LALPUR PATTI KHURD</option>
                        <option>LALPUR PATTI KUNDAN</option>
                        <option>LALPUR PATTI QADAR KHAN</option>
                        <option>LALPUR PATTI SAEED GANJ</option>
                        <option>LALUNAGLA-1</option>
                        <option>LALUNAGLA2</option>
                        <option>LALWARA -1</option>
                        <option>LALWARA 02</option>
                        <option>LAMBAKHEDA-1</option>
                        <option>LAMBAKHEDA-2</option>
                        <option>LAMBAKHEDA-3</option>
                        <option>LAMBAKHEDA-4</option>
                        <option>LASHKAR GANJ</option>
                        <option>LEHTORA</option>
                        <option>LODHIPUR</option>
                        <option>LODHIPUR GULADIYA</option>
                        <option>LODHIPUR NAYAK-1</option>
                        <option>LODHIPUR NAYAK-2</option>
                        <option>LOHAPATTI BHAGIRATH</option>
                        <option>LOHAPATTI BHOLANATH-1</option>
                        <option>LOHAPATTI BHOLANATH-2</option>
                        <option>LOHARA TANDA</option>
                        <option>LOHARI</option>
                        <option>LOHARRA INAYATGANJ</option>
                        <option>MADAIRYAN KALI</option>
                        <option>MADAIYAN BHAJJAN</option>
                        <option>MADAIYAN JAULAPUR</option>
                        <option>MADAIYAN TULSI</option>
                        <option>MADAIYAN VADE</option>
                        <option>MADANAGLA</option>
                        <option>MADARPUR</option>
                        <option>MADEYAM DHOLSAR</option>
                        <option>MADEYAN ODEYRAJ</option>
                        <option>MADEYAN SENDOALI</option>
                        <option>MADHUKAR 01</option>
                        <option>MADHUKAR 02</option>
                        <option>MADHUKAR-3</option>
                        <option>MADHUPARI -01</option>
                        <option>MADHUPARI -02</option>
                        <option>MADHUPURA-1</option>
                        <option>MADHUPURA-2</option>
                        <option>MADIYAN JHAU</option>
                        <option>MADRASA ALIYA</option>
                        <option>MADRASA KOHANA</option>
                        <option>MAGARMAU</option>
                        <option>MAHAMUD PUR</option>
                        <option>MAHARAJ PUR</option>
                        <option>MAHESHPURA</option>
                        <option>MAHEVA</option>
                        <option>MAHTOSH-1</option>
                        <option>MAHTOSH-2</option>
                        <option>MAHUA KHERA SUAR</option>
                        <option>MAHUAKHEDA TANDA</option>
                        <option>MAHUNAGAR -1</option>
                        <option>MAHUNAGAR -2</option>
                        <option>MAHUNAGAR -3</option>
                        <option>MAHUNAGAR 01</option>
                        <option>MAHUNAGAR 02</option>
                        <option>MAINI</option>
                        <option>MAJAR BAGADADI SAHAB</option>
                        <option>MAJAR CHUPASHAH MIYA</option>
                        <option>MAJAR SHAH DARAGAHI SAHAB</option>
                        <option>MAJAR VAJIR ALI SAHAB</option>
                        <option>MAJHRA</option>
                        <option>MAJHRA  KALINAGAR  LACHHIWALA MAJHRA KHEMWALA -1</option>
                        <option>MAJHRA  KALINAGAR  LACHHIWALA MAJHRA KHEMWALA -2</option>
                        <option>MAJHRA AMEER KHAN</option>
                        <option>MAJULLANAGAR-1</option>
                        <option>MAJULLANAGAR-2</option>
                        <option>MAJULLANAGAR-3</option>
                        <option>MAJULLANAGAR-4</option>
                        <option>MALANKHEDA</option>
                        <option>MALGOSHA</option>
                        <option>MALHAPURA</option>
                        <option>MALKI KAMALPUR</option>
                        <option>MAMURPUR BADOULI</option>
                        <option>MANAKPUR ANZARIYA</option>
                        <option>MANAONA</option>
                        <option>MANDBA HASANPUR</option>
                        <option>MANDBA HASANPUR-2</option>
                        <option>MANDEYAN BALLU</option>
                        <option>MANDEYAN RAMI</option>
                        <option>MANDEYAN SHADI-1</option>
                        <option>MANDEYAN SHADI-2</option>
                        <option>MANDHOLI-1</option>
                        <option>MANDOLI-2</option>
                        <option>MANGADPUR</option>
                        <option>MANGOLI 01</option>
                        <option>MANGOLI 02</option>
                        <option>MANGUPURA</option>
                        <option>MANIHAR KHEDA</option>
                        <option>MANKARA</option>
                        <option>MANPUR CHHAPAT</option>
                        <option>MANPUR OJHA-1</option>
                        <option>MANPUR OJHA-2</option>
                        <option>MANPUR OJHA-3</option>
                        <option>MANPUR OJHA-4</option>
                        <option>MANPUR UTTARI-1</option>
                        <option>MANPUR UTTARI-2</option>
                        <option>MANSOORPUR</option>
                        <option>MANSUR PUR</option>
                        <option>MANSURPUR-1</option>
                        <option>MANSURPUR-2</option>
                        <option>MANULLAPUR</option>
                        <option>MANUNAGAR</option>
                        <option>MARGHATI</option>
                        <option>MASJID DAROGA MAHABUB JAN</option>
                        <option>MASJID JASAUDI PRADHAN</option>
                        <option>MASJID KAITHAVALI</option>
                        <option>MASJID KALE KHAN</option>
                        <option>MASJID KALLASH MAY GUIYA TALAB-1</option>
                        <option>MASJID KALLASH MAY GUIYA TALAB-2</option>
                        <option>MASJID MOLANA JAHURUDDIN  KHAN-1</option>
                        <option>MASJID MOLANA JAHURUDDIN KHAN 2</option>
                        <option>MASJID QAZI</option>
                        <option>MASWASI 5</option>
                        <option>MASWASI-1</option>
                        <option>MASWASI-2</option>
                        <option>MASWASI-3</option>
                        <option>MASWASI-4</option>
                        <option>MATHURAPUR -2</option>
                        <option>MATHURAPUR -3</option>
                        <option>MATHURAPUR 01</option>
                        <option>MATKHEDA-1</option>
                        <option>MATKHEDA-2</option>
                        <option>MATWALI --1</option>
                        <option>MATWALI -2</option>
                        <option>MAU</option>
                        <option>MAU. BARADRI MAHAMUD KHAN</option>
                        <option>MAZRA TRILOK PUR</option>
                        <option>MDEAYAN GOR</option>
                        <option>MDEYAN BUDHPUR</option>
                        <option>MEERAPUR</option>
                        <option>MEERAPUR MEERGANJ-1</option>
                        <option>MEGHANAGLA JADID</option>
                        <option>MEGHANAGLA KADIM</option>
                        <option>MEHANDI NAGAR JUNUBI</option>
                        <option>MEHEL KUMEDAN</option>
                        <option>MEHNDINAGAR</option>
                        <option>MEHNDINAGAR SHUMALI</option>
                        <option>MEHNDIPUR</option>
                        <option>MENDIPUR</option>
                        <option>MENHDINAGAR</option>
                        <option>MEWALA DHARU</option>
                        <option>MEWLA KALA</option>
                        <option>MIAN GANJ</option>
                        <option>MILAK  SANYYA</option>
                        <option>MILAK AARIF</option>
                        <option>MILAK ABBU KHAN</option>
                        <option>MILAK ASAD KHAN</option>
                        <option>MILAK ASDULLAPUR-1</option>
                        <option>MILAK ASDULLAPUR-2</option>
                        <option>MILAK ASDULLAPUR-3</option>
                        <option>MILAK ASDULLAPUR-4</option>
                        <option>MILAK ASDULLAPUR-5</option>
                        <option>MILAK BANKERALI</option>
                        <option>MILAK BHURE KHAN</option>
                        <option>MILAK DUNDI</option>
                        <option>MILAK DYUDI INAYAT RASOOL KHAN</option>
                        <option>MILAK FEJULLANAGAR</option>
                        <option>MILAK GULAM ALI</option>
                        <option>MILAK GULAM KHAN</option>
                        <option>MILAK HAFIZ BALLI</option>
                        <option>MILAK HASHAM</option>
                        <option>MILAK HASHAM IBRAHEEM KHAN</option>
                        <option>MILAK ICHHCARAM</option>
                        <option>MILAK KHANAM</option>
                        <option>MILAK KHANPUR</option>
                        <option>MILAK KHOOD</option>
                        <option>MILAK MAFI PAIPURA</option>
                        <option>MILAK MAJBOOTA</option>
                        <option>MILAK MALANG GAAN</option>
                        <option>MILAK MIRZA FAYYAZ</option>
                        <option>MILAK MOHAMMAD BAKSH</option>
                        <option>MILAK MOHMMAD NAKI</option>
                        <option>MILAK MOHMMAD SHAH KHAN</option>
                        <option>MILAK MUGLA URF MUGLPIUR</option>
                        <option>MILAK MUNDI</option>
                        <option>MILAK NAGALI</option>
                        <option>MILAK NASIRABAD-1</option>
                        <option>MILAK NASIRABAD-2</option>
                        <option>MILAK NASIRABAD-3</option>
                        <option>MILAK NIBBI SINGH</option>
                        <option>MILAK NOUKHAREED</option>
                        <option>MILAK PIPALSANA</option>
                        <option>MILAK QAZI</option>
                        <option>MILAK SHADI NAGAR</option>
                        <option>MILAK SIKARAUL</option>
                        <option>MILAK TAHABBAR ALI KHAN</option>
                        <option>MILAK TAJ KHAN</option>
                        <option>MILAK TURKHEDA</option>
                        <option>MILAK WADULLA</option>
                        <option>MILAK YAKOOB ALI KHAN</option>
                        <option>MILAL AFZAL KHAN</option>
                        <option>MILETRI LAIN-1</option>
                        <option>MILETRI LAIN-2</option>
                        <option>MIRAPUR BILASPUR</option>
                        <option>MIRAPUR MEERGANJ -2</option>
                        <option>MIRJAPUR CHAKARPUR</option>
                        <option>MIRJAPUR PARAM</option>
                        <option>MIRZAPUR BILASPUR</option>
                        <option>MIRZAPUR SUAR</option>
                        <option>MISHRINAGAR JUNUBI</option>
                        <option>MISRINAGAR SHUMALI</option>
                        <option>MISTAN GANJ AKAB MISTAN GANJ</option>
                        <option>MITTARPUR AHROULA 01</option>
                        <option>MITTARPUR AHROULA 02</option>
                        <option>MOHABBAT NAGAR KUNDESRA</option>
                        <option>MOHABBATNAGAR NEAR SUAR</option>
                        <option>MOHALIYA</option>
                        <option>MOHAMMAD NAGAR</option>
                        <option>MOHAMMAD NAGAR KHUTIA</option>
                        <option>MOHAMMADGUNJ</option>
                        <option>MOHAMMADNAGAR</option>
                        <option>MOHAMMADNAGAR NANKAR</option>
                        <option>MOHAMMADPUR JADID</option>
                        <option>MOHAMMADPUR KADIM</option>
                        <option>MOHANPUR</option>
                        <option>MOHBBATGANJ TANDA</option>
                        <option>MOHD PUR SHUMALI</option>
                        <option>MOLANA MOHAMMAD ALI JOHAR MARG MAY GANGAPUR HATS-9</option>
                        <option>MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HARTS 8</option>
                        <option>MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HATS-7</option>
                        <option>MOMINPUR AHAMDABAD-1</option>
                        <option>MOMINPUR AHAMDABAD-2</option>
                        <option>MORI GET-1</option>
                        <option>MORI GET-2</option>
                        <option>MORI GET-3</option>
                        <option>MOTIPURA</option>
                        <option>MUBANA</option>
                        <option>MUBARAKPUR</option>
                        <option>MUBARAKPUR-1</option>
                        <option>MUBARAKPUR-2</option>
                        <option>MUDI</option>
                        <option>MUDIYA KALAN</option>
                        <option>MUDIYA KHURD</option>
                        <option>MUDIYA NEAR MANKARA</option>
                        <option>MUHAMMADPUR BAS KHADIYA</option>
                        <option>MUJHIYANA</option>
                        <option>MUKRABPUR</option>
                        <option>MUKRMPUR</option>
                        <option>MUKUTPUR</option>
                        <option>MULLAKHERA</option>
                        <option>MUNDIA KHEDA</option>
                        <option>MUNDIYA</option>
                        <option>MUNDIYA KALAN</option>
                        <option>MUNDIYA RASOOLPUR</option>
                        <option>MUNDIYA SEDNAGAR</option>
                        <option>MURSAINA</option>
                        <option>MUSARFGANJ</option>
                        <option>MUSTAFABAD DAUNKPURI</option>
                        <option>MUSTAFABAD KHURD</option>
                        <option>MUSTAFABAD KLA</option>
                        <option>MUSTAFABAD URF TAKLABAD</option>
                        <option>MUTIYAPURA BAZARPATTI</option>
                        <option>MUTIYAPURA NIKAT BHOT</option>
                        <option>NAANKAR</option>
                        <option>NAANKAR RANI-1</option>
                        <option>NAANKAR RANI-2</option>
                        <option>NAANKAR RANI-3</option>
                        <option>NAANKAR RANI-4</option>
                        <option>NABAB  NAGAR TAH</option>
                        <option>NABAB NAGAR TANDA</option>
                        <option>NABABGANJ SHUMALI</option>
                        <option>NABABGUNJ</option>
                        <option>NABBA NAGLA</option>
                        <option>NABI GANJ PIPLI</option>
                        <option>NABIGANJ</option>
                        <option>NABIGANJ JADID</option>
                        <option>NADARGANJ</option>
                        <option>NADNA</option>
                        <option>NADNAU</option>
                        <option>NAEEM GANJ</option>
                        <option>NAGALIYA</option>
                        <option>NAGALIYA KASAMGANJ</option>
                        <option>NAGARIYA - 2</option>
                        <option>NAGARIYA KALAN -1</option>
                        <option>NAGARIYA KALAN -2</option>
                        <option>NAGARIYA-1</option>
                        <option>NAGLA BANSNAGLI</option>
                        <option>NAGLA GANESH</option>
                        <option>NAGLA UDAI-1</option>
                        <option>NAGLA UDAI-2</option>
                        <option>NAGLI BHAGWANT</option>
                        <option>NAGLIYA AAQIL 1</option>
                        <option>NAGLIYA AAQIL-2</option>
                        <option>NAGLIYA AAQIL-4</option>
                        <option>NAGLIYA AAQIL-5</option>
                        <option>NAGLIYA AAQIL-6</option>
                        <option>NAGLIYA AAQIL-7</option>
                        <option>NAGLIYA AAQIL-8</option>
                        <option>NAGRIYA KHURD</option>
                        <option>NAIMGANJ</option>
                        <option>NALAPAR</option>
                        <option>NALAPAR MAY KUCHA SAUDAGARAN-2</option>
                        <option>NALAPAR MAY KUCHASAUDAGARAN-1</option>
                        <option>NANDGAUN</option>
                        <option>NANKAR</option>
                        <option>NANKAREN</option>
                        <option>NARAYAN NAGALA</option>
                        <option>NARAYANPUR</option>
                        <option>NARKHEDA</option>
                        <option>NARKHEDA-1</option>
                        <option>NARKHEDA-2</option>
                        <option>NARKHEDI</option>
                        <option>NARKHERA</option>
                        <option>NARKHERI</option>
                        <option>NARPATNAGAR-1</option>
                        <option>NARPATNAGAR-2</option>
                        <option>NARPATNAGAR-3</option>
                        <option>NARPATNAGAR-4</option>
                        <option>NARPATNAGAR-5</option>
                        <option>NARPATNAGAR-6</option>
                        <option>NARPATNAGAR-7</option>
                        <option>NARPATNAGAR-8</option>
                        <option>NARSUA</option>
                        <option>NASIMGANJ</option>
                        <option>NASRAT NAGAR</option>
                        <option>NASRATNAGAR</option>
                        <option>NAUGAVAN-1</option>
                        <option>NAUGAVAN-2</option>
                        <option>NAUGJA</option>
                        <option>NAVAB GANJ PARSAL</option>
                        <option>NAVAB NAGAR PADPURI</option>
                        <option>NAVADIYA</option>
                        <option>NAVAVPURA</option>
                        <option>NAVI GANJ BATHUAKHEDA</option>
                        <option>NAVIGUJ KADEEM</option>
                        <option>NAWAB GAUNJ JUNUBI</option>
                        <option>NAWABGANJ</option>
                        <option>NAWABNAGAR NEAR PATWAI</option>
                        <option>NAWADA</option>
                        <option>NAYA GANAO</option>
                        <option>NAYA GANV NAZIBABAD</option>
                        <option>NAYA GAON</option>
                        <option>NEW SAGARPUR, BISOULI -01</option>
                        <option>NEW SAGARPUR, BISOULI -02</option>
                        <option>NEW SAGARPUR, BISOULI -03</option>
                        <option>NEWVILLAGE NIKAT AKBARABAD</option>
                        <option>NIPANIYA-1</option>
                        <option>NIPANIYA-2</option>
                        <option>NISHVI</option>
                        <option>NISVI</option>
                        <option>NISWA</option>
                        <option>NIYAMAT NAGAR</option>
                        <option>NOORPUR HAZOORPUR</option>
                        <option>NORAHA</option>
                        <option>NORANGPUR</option>
                        <option>NRENDARPUR 01</option>
                        <option>NRENDARPUR 02</option>
                        <option>NURPUR</option>
                        <option>NYU AVAS VIKAS</option>
                        <option>ORANGNAGAR KHEDA</option>
                        <option>OSI 01</option>
                        <option>OSI 02</option>
                        <option>PACHTOUR</option>
                        <option>PADMAPUR</option>
                        <option>PADPURI</option>
                        <option>PADPURI NEAR KALYANPUR</option>
                        <option>PADPURI SUAR</option>
                        <option>PAGAMBERPUR</option>
                        <option>PAHADI</option>
                        <option>PAHADI GET</option>
                        <option>PAHADPUR</option>
                        <option>PAHADPUR BILASPUR</option>
                        <option>PAIGAMBAR PUR</option>
                        <option>PAIGAMBARPUR</option>
                        <option>PAIGUPURA</option>
                        <option>PAIMPUR</option>
                        <option>PAINDA NAGAR</option>
                        <option>PAIPURA</option>
                        <option>PAJABA</option>
                        <option>PAJIYA</option>
                        <option>PAJWA-1</option>
                        <option>PAJWA-2</option>
                        <option>PALPURA</option>
                        <option>PANJABIYAN</option>
                        <option>PANJABNAGAR-1</option>
                        <option>PANJABNAGAR-2</option>
                        <option>PANJABNAGAR-3</option>
                        <option>PANJAWNAGAR</option>
                        <option>PANVADIYA</option>
                        <option>PANVADIYA KVATARS 2</option>
                        <option>PANVADIYA KVATARS-1</option>
                        <option>PANWADIYA</option>
                        <option>PARAM-1</option>
                        <option>PARAM-2</option>
                        <option>PARAM-3</option>
                        <option>PARAM-4</option>
                        <option>PARBATPUR</option>
                        <option>PARCHAI</option>
                        <option>PAREWA</option>
                        <option>PAROUTA 01</option>
                        <option>PARSHUPURA</option>
                        <option>PARSIPURA</option>
                        <option>PARTAPUR HARDASPUR</option>
                        <option>PASHUPURA</option>
                        <option>PASIYAPURA</option>
                        <option>PASIYAPURA JANOOBI</option>
                        <option>PASIYAPURA SHUMALI</option>
                        <option>PASUPURA</option>
                        <option>PATARIYA</option>
                        <option>PATIYA</option>
                        <option>PATLE V MASJID KHATOTI</option>
                        <option>PATNA</option>
                        <option>PATTHAR KHEDA</option>
                        <option>PATTI ASHOKPUR</option>
                        <option>PATTI BASANTPUR</option>
                        <option>PATTI FAZILABAD</option>
                        <option>PATTI KALAN-2</option>
                        <option>PATTI KALYANPUR</option>
                        <option>PATTI TOLA</option>
                        <option>PATTIKLA-1</option>
                        <option>PATWAI -1</option>
                        <option>PATWAI -2</option>
                        <option>PATWAI -3</option>
                        <option>PATWAI 04</option>
                        <option>PEGA</option>
                        <option>PEGAMWARPUR</option>
                        <option>PESH KOTAVALI</option>
                        <option>PILA TALAB</option>
                        <option>PIPAL TOLA</option>
                        <option>PIPALGAUN-1</option>
                        <option>PIPALGAUN-2</option>
                        <option>PIPALIYA AHALA</option>
                        <option>PIPALIYA BHOJ</option>
                        <option>PIPALIYA GOPAL</option>
                        <option>PIPALIYA GOPAL-2</option>
                        <option>PIPALIYA JAT</option>
                        <option>PIPALIYA MEHTO</option>
                        <option>PIPALIYA MISRA</option>
                        <option>PIPALIYA NAV</option>
                        <option>PIPALIYA RAIJADA</option>
                        <option>PIPALIYA VIJAYNAGAR</option>
                        <option>PIPALSANA</option>
                        <option>PIPALWALA</option>
                        <option>PIPLA SHIV NAGAR-2</option>
                        <option>PIPLA SHIVNAGAR-1</option>
                        <option>PIPLATOLA NIKAT GUJRATOLA 2</option>
                        <option>PIPLATOLA NIKAT GUJRATOLA-1</option>
                        <option>PIPLI</option>
                        <option>PIPLI NAYAK-1</option>
                        <option>PIPLI NAYAK-2</option>
                        <option>PIPLI NAYAK-3</option>
                        <option>POHAPPUR</option>
                        <option>PRANPUR</option>
                        <option>PRITHVI NAGAR</option>
                        <option>PRITHVIPUR &amp; CHIDIYAKHEDA</option>
                        <option>PROTA -2</option>
                        <option>PSIYAPURA</option>
                        <option>PUL PUKHTA</option>
                        <option>PULIS LAIN</option>
                        <option>PURENA</option>
                        <option>PURENIYA JADID</option>
                        <option>PURENIYA KALAN</option>
                        <option>PURENIYA KHURD</option>
                        <option>PURNAPUR</option>
                        <option>PUSWADA-1</option>
                        <option>PUSWADA-2</option>
                        <option>RABNNA</option>
                        <option>RAFATPUR</option>
                        <option>RAFTPUR</option>
                        <option>RAIKANAGALA</option>
                        <option>RAIPUR</option>
                        <option>RAIPUR-01</option>
                        <option>RAIPUR-02</option>
                        <option>RAIPUR-1</option>
                        <option>RAIPUR-2</option>
                        <option>RAJA RAMPUR</option>
                        <option>RAJADARA NAVAB GET V MASJID PILU-1</option>
                        <option>RAJADARA NAVAB GET V MASJID PILU-2</option>
                        <option>RAJADARA NAVAB GET V MASJID PILU-3</option>
                        <option>RAJADARA NAVAB GET V MASJID PILU-4</option>
                        <option>RAJADARA NAVAB GET V MASJID PILU-5</option>
                        <option>RAJJAD</option>
                        <option>RAJOURA</option>
                        <option>RAJPURA</option>
                        <option>RAJPURA SUAR</option>
                        <option>RAJPURA TANDA</option>
                        <option>RAMNAGAR</option>
                        <option>RAMNAGAR LATIFPUR-1</option>
                        <option>RAMNAGAR LATIFPUR-2</option>
                        <option>RAMNAGAR LATIFPUR-3</option>
                        <option>RAMNAGARIYA</option>
                        <option>RAMNAGRIYA</option>
                        <option>RAMNATH KALONI MAY GOVAN ASPATAL</option>
                        <option>RAMPUR DHAMMAN</option>
                        <option>RAMPURA</option>
                        <option>RAMPURA BUJURGA-1</option>
                        <option>RAMPURA BUJURGA-2</option>
                        <option>RANUA NAGALA</option>
                        <option>RASDANDIYA</option>
                        <option>RASOOLPUR</option>
                        <option>RASOOLPUR  -1</option>
                        <option>RASOOLPUR  -2</option>
                        <option>RASOOLPUR  -3</option>
                        <option>RASOOLPUR  -4</option>
                        <option>RASOOLPUR  -5</option>
                        <option>RASOOLPUR  -6</option>
                        <option>RASOOLPUR FAREEDPUR-1</option>
                        <option>RASOOLPUR FAREEDPUR-2</option>
                        <option>RASULA</option>
                        <option>RASULPUR</option>
                        <option>RATANPURA</option>
                        <option>RATANPURA SHUMALI 1</option>
                        <option>RATANPURA SHUMALI 2</option>
                        <option>RATANPURA-1</option>
                        <option>RATANPURA-2</option>
                        <option>RATHONDA-1</option>
                        <option>RATHONDA-2</option>
                        <option>RATUANAGALA-1</option>
                        <option>RATUANAGALA-2</option>
                        <option>RATUANAGALA-3</option>
                        <option>RAVANA -01</option>
                        <option>RAVANA -02</option>
                        <option>RAVANAGHASI</option>
                        <option>RAWANALALA &amp; VISHANPURI</option>
                        <option>RAWANIPATTI JBHAHAR</option>
                        <option>RAWANIPATTI UDA</option>
                        <option>RAYPUR -1</option>
                        <option>RAYPUR-2</option>
                        <option>RAYPUR-3</option>
                        <option>RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 1</option>
                        <option>RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 2</option>
                        <option>RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 3</option>
                        <option>RAZA NAGAR-1</option>
                        <option>RAZA NAGAR-2</option>
                        <option>REHAT GANJ</option>
                        <option>REHPURA</option>
                        <option>REHSENA-1</option>
                        <option>REHSENA-2</option>
                        <option>RELVE STESH N ERIYA</option>
                        <option>REVERI KALAN 01</option>
                        <option>REVERI KALAN 02</option>
                        <option>RIVDI KHURD</option>
                        <option>ROARA KALAN-1</option>
                        <option>ROARA KALAN-2</option>
                        <option>ROARA KHURD-1</option>
                        <option>ROARA KHURD-2</option>
                        <option>ROARA KHURD-3</option>
                        <option>ROOPPUR</option>
                        <option>RUPAPUR</option>
                        <option>RUSTAMNAGAR</option>
                        <option>RUSTAMNAGAR NEAR CHAPRRA</option>
                        <option>RUSTAMNAGAR NEAR RAYPUR</option>
                        <option>RUSTAMPUR</option>
                        <option>RUTAMPUR</option>
                        <option>SADAR SAFAKHANA -3</option>
                        <option>SADAR SAFAKHANA-1</option>
                        <option>SADAR SAFAKHANA-2</option>
                        <option>SADARAKHEDA</option>
                        <option>SADAT 01</option>
                        <option>SADAT 02</option>
                        <option>SADDIKNAGAR</option>
                        <option>SAEDABAD</option>
                        <option>SAEED NAGAR</option>
                        <option>SAHARIYA JAWHAR</option>
                        <option>SAHAVIAN KHURD</option>
                        <option>SAHAVIYAN KALAN</option>
                        <option>SAHRIYA DRAJ</option>
                        <option>SAHRIYA LAKKKHA</option>
                        <option>SAHRIYA NARPAT</option>
                        <option>SAHUKARA-1</option>
                        <option>SAHUKARA-2</option>
                        <option>SAHUKARA-3</option>
                        <option>SAHUKARA-4</option>
                        <option>SAHUKARA-5</option>
                        <option>SAHUKARA-6</option>
                        <option>SAHUNAGLA</option>
                        <option>SAID ALIGANJ BANZARIYA</option>
                        <option>SAID NAGAR TANDA</option>
                        <option>SAIDNAGAR BAZARPTTI</option>
                        <option>SAIDNAGAR JANUBI</option>
                        <option>SAIDNAGAR URF MADAIYAN PUSE</option>
                        <option>SAIFNI -01</option>
                        <option>SAIFNI -02</option>
                        <option>SAIFNI -03</option>
                        <option>SAIFNI -04</option>
                        <option>SAIFNI -05</option>
                        <option>SAIFNI -06</option>
                        <option>SAIFNI -07</option>
                        <option>SAIFNI -08</option>
                        <option>SAIFNI -09</option>
                        <option>SAIFNI -10</option>
                        <option>SAIFNI -11</option>
                        <option>SAKATPURA</option>
                        <option>SAKATUBA</option>
                        <option>SALABATPUR</option>
                        <option>SALVE NAGAR</option>
                        <option>SAMODIYA-1</option>
                        <option>SAMODIYA-2</option>
                        <option>SANAYYA CHAK</option>
                        <option>SANAYYA JATT</option>
                        <option>SANAYYA SUKH</option>
                        <option>SANIYA AHMAD</option>
                        <option>SANKARA</option>
                        <option>SARABA-1</option>
                        <option>SARABA-2</option>
                        <option>SARAI IMAM</option>
                        <option>SARAY DARAVAJA-1</option>
                        <option>SARAY DARAVAJA-2</option>
                        <option>SARAY KADEEM</option>
                        <option>SARAY KALAN-1</option>
                        <option>SARAY KALAN-2</option>
                        <option>SARAY MAHESH -1</option>
                        <option>SARAY MAHESH -2</option>
                        <option>SARAY SAHADAT YAR KHAN-1</option>
                        <option>SARAY SAHADAT YAR KHAN-2</option>
                        <option>SARAYE JAHANGIR -1</option>
                        <option>SARAYE JAHANGIR -2</option>
                        <option>SARAYE PUKHTA</option>
                        <option>SARBARNAGAR</option>
                        <option>SARDHARPUR</option>
                        <option>SARKADA</option>
                        <option>SARKADI</option>
                        <option>SARKADI -1</option>
                        <option>SARKADI -2</option>
                        <option>SARTHAL</option>
                        <option>SATUNE SANG -1</option>
                        <option>SATUNE SANG -2</option>
                        <option>SATUNE SANG -3</option>
                        <option>SATUNE SANG -4</option>
                        <option>SATUNE SANG -5</option>
                        <option>SED NAGAR NEAR ASALATPUR</option>
                        <option>SEDABAD</option>
                        <option>SEDNAGAR MUNDIYA</option>
                        <option>SEDU KA MAZHARA</option>
                        <option>SEDULLANAGAR</option>
                        <option>SEJNINANKAR-1</option>
                        <option>SEJNINANKAR-2</option>
                        <option>SEJNINANKAR-3</option>
                        <option>SEJNINANKAR-4</option>
                        <option>SEJNINANKAR-5</option>
                        <option>SEJNINANKAR-6</option>
                        <option>SEMRA NEAR LADPUR-1</option>
                        <option>SEMRA NEAR LADPUR-2</option>
                        <option>SENDOALI</option>
                        <option>SENJNA</option>
                        <option>SENPUR</option>
                        <option>SENTAKHEDA-1</option>
                        <option>SENTAKHEDA-2</option>
                        <option>SENTAKHEDA-3</option>
                        <option>SENTAKHEDA-4</option>
                        <option>SENTAKHEDA-5</option>
                        <option>SENTAKHEDA-6</option>
                        <option>SHADARAUFR THUNAPUR</option>
                        <option>SHADI NAGAR HAJIRA-2</option>
                        <option>SHADINAGAR</option>
                        <option>SHADINAGAR HARDASPUR</option>
                        <option>SHADINAGAR HAZIRA 1</option>
                        <option>SHADINAGAR SAKATPURA</option>
                        <option>SHADINGAR NI0AHAMADNAGAR JAGIR</option>
                        <option>SHAFAYAT NAGAR</option>
                        <option>SHAHABAD GET-1</option>
                        <option>SHAHABAD GET-2</option>
                        <option>SHAHAJAD NAGAR</option>
                        <option>SHAHDARA</option>
                        <option>SHAHDRA NEAR  BHOT</option>
                        <option>SHAHDRA NIKAT DHANPUR</option>
                        <option>SHAHPUR DEV 01</option>
                        <option>SHAHPUR DEV 02</option>
                        <option>SHAHPURA-1</option>
                        <option>SHAHPURA-2</option>
                        <option>SHAHPURA-3</option>
                        <option>SHAMSHABAD</option>
                        <option>SHAMSHABAD  URF KHABRIYA</option>
                        <option>SHAZAD NAGAR</option>
                        <option>SHEHZADNAGAR</option>
                        <option>SHEKHUPURA</option>
                        <option>SHIKARPUR</option>
                        <option>SHIKARUR</option>
                        <option>SHIV NAGAR NEAR ASALATPUR</option>
                        <option>SHIV NAGAR NEAR GULADIYA</option>
                        <option>SHIV PURI</option>
                        <option>SHIV VIHAR</option>
                        <option>SHIVNAGAR NEAR LOHARI</option>
                        <option>SHIVNAGAR NEAR RAMNAGAR</option>
                        <option>SHIVNAGAR NIKAT GANGAPUR</option>
                        <option>SHOKAT ALI MARG -1</option>
                        <option>SHOKAT ALI MARG -2</option>
                        <option>SHOKAT ALI MARG -3</option>
                        <option>SHOKAT ALI MARG -4</option>
                        <option>SHOKAT ALI MARG -5</option>
                        <option>SHRI NAGAR</option>
                        <option>SHUTRAKHANA MAY JYOTISHYAN-1</option>
                        <option>SHUTRAKHANA MAY JYOTISHYAN-2</option>
                        <option>SHUTRAKHANA MAY JYOTISHYAN-3</option>
                        <option>SHYAMPUR</option>
                        <option>SIGANI</option>
                        <option>SIGANKHEDA-1</option>
                        <option>SIGANKHEDA-2</option>
                        <option>SIGANKHEDA-3</option>
                        <option>SIGANKHEDA-4</option>
                        <option>SIGANKHEDA-5</option>
                        <option>SIGANKHEDA-6</option>
                        <option>SIGANKHEDA-7</option>
                        <option>SIHARI-1</option>
                        <option>SIHARI-2</option>
                        <option>SIHARIYA</option>
                        <option>SIHOR</option>
                        <option>SIHORA</option>
                        <option>SIKANDARABAD-1</option>
                        <option>SIKANDARABAD-2</option>
                        <option>SIKARAUL</option>
                        <option>SIKHADIYAN -1</option>
                        <option>SIKHADIYAN -2</option>
                        <option>SIKHADIYAN -3</option>
                        <option>SIKHADIYAN -4</option>
                        <option>SIKRAURA</option>
                        <option>SILAI BARA GAON-1</option>
                        <option>SILAI BARA GAON-2</option>
                        <option>SIMARIYA-1</option>
                        <option>SIMARIYA-2</option>
                        <option>SIMRA</option>
                        <option>SINGRA-1</option>
                        <option>SINGRA-2</option>
                        <option>SIRASKHERA</option>
                        <option>SIRKAPATTI KUNDESARI</option>
                        <option>SIRKAPATTI RUPAPUR</option>
                        <option>SIRRA</option>
                        <option>SIRSA</option>
                        <option>SISAUNA-1</option>
                        <option>SISAUNA-2</option>
                        <option>SITAURA</option>
                        <option>SITLA</option>
                        <option>SOHANA</option>
                        <option>SONAKPUR-1</option>
                        <option>SONAKPUR-2</option>
                        <option>SONAKPUR-3</option>
                        <option>SOOPA</option>
                        <option>STATION ROAD-1</option>
                        <option>STATION ROAD-2</option>
                        <option>STATION ROAD-3 MAY GOVAN COLONY</option>
                        <option>SUANAGLA</option>
                        <option>SUAR KHAS DAKSHNI -1</option>
                        <option>SUAR KHAS DAKSHNI -2</option>
                        <option>SUAR KHAS DAKSHNI -3</option>
                        <option>SUAR KHAS UTTARI</option>
                        <option>SUAR KHURD</option>
                        <option>SUHAG NAGLA</option>
                        <option>SUHAVA</option>
                        <option>SUJANPUR</option>
                        <option>SUJATGANJ KHERA</option>
                        <option>SUJAYATGUNJ</option>
                        <option>SUNARKHEDA</option>
                        <option>SURAJPUR</option>
                        <option>SURATSINGHPUR URF NEW VILLAGE-1</option>
                        <option>SURATSINGHPUR URF NEW VILLAGE-2</option>
                        <option>SURAZPUR</option>
                        <option>SURJAPUR</option>
                        <option>SWALEPUR</option>
                        <option>SWAR KHURD-1</option>
                        <option>SWAR KHURD-2</option>
                        <option>SYONDRA</option>
                        <option>TAH KHURD</option>
                        <option>TAHKLA</option>
                        <option>TAJPUR</option>
                        <option>TAJPUR BEHTA</option>
                        <option>TAJPUR LAKHAN</option>
                        <option>TAKIYA -1</option>
                        <option>TAKIYA FATEH SHAH KHAN</option>
                        <option>TAKIYA MEMARAN</option>
                        <option>TAKIYA MUBARAK SHAH KHA</option>
                        <option>TAKIYA SARVAR SHAH KHAN</option>
                        <option>TAKIYA-2</option>
                        <option>TALAB CHAMARAN</option>
                        <option>TALAB KELE VALA</option>
                        <option>TALAB MAJAR BAGADADI SAHAB</option>
                        <option>TALAB MULLA AIRAM</option>
                        <option>TALAB NIHALUDDIN MAY GHER HASSAN KHAN</option>
                        <option>TALABPUR-1</option>
                        <option>TALABPUR-2</option>
                        <option>TALIBABAD</option>
                        <option>TALMAHAWAR</option>
                        <option>TANDA 01</option>
                        <option>TANDA 02</option>
                        <option>TANDA GOLU</option>
                        <option>TARAUA</option>
                        <option>TASHKA -1</option>
                        <option>TASHKA- 2</option>
                        <option>TATATVALAN MAY MAJAR TATASHAH MIYAN</option>
                        <option>TEHRI KHAWJA-1</option>
                        <option>TEHRI KHAWJA-2</option>
                        <option>TELIPURA</option>
                        <option>TEMRA</option>
                        <option>THAKURDUWARA</option>
                        <option>THANATIN-1</option>
                        <option>THANATIN.2</option>
                        <option>THANATIN.3</option>
                        <option>THANATIN.4</option>
                        <option>THANATIN.5</option>
                        <option>THANATIN.6</option>
                        <option>THEGA</option>
                        <option>THIRIA VISHNU</option>
                        <option>THIRIYA JADID</option>
                        <option>THIRIYA SALEPUR</option>
                        <option>THOTHAR MAY SABJI MANDI -2</option>
                        <option>THOTHAR MAY SABZI MANDI - 1</option>
                        <option>TICKET GANJ</option>
                        <option>TIRAH</option>
                        <option>TODIPURA</option>
                        <option>TOPAKHANA ROAD</option>
                        <option>TUMADIYA</option>
                        <option>TUMADIYA-1</option>
                        <option>TUMADIYA-2</option>
                        <option>TURKHERA</option>
                        <option>UDAIPUR JAGIR</option>
                        <option>UDAYEPUR</option>
                        <option>UDAYPUR JAGIR</option>
                        <option>UDHAMPUR</option>
                        <option>UMRI</option>
                        <option>UNCHA GAUN 01</option>
                        <option>UNCHA GAUN 02</option>
                        <option>VAMNAPURI GET-1</option>
                        <option>VAMNAPURI GET-2</option>
                        <option>VEERPUR</option>
                        <option>VENAJIRPURA  URF GHATAMPUR 2</option>
                        <option>VENAJIRPURA URF GHATAMPUR-1</option>
                        <option>VIJIYA</option>
                        <option>VIKRAMPUR</option>
                        <option>VISHANPUR JAGEER</option>
                        <option>VISHARATNAGAR-1</option>
                        <option>VISHARATNAGAR-2</option>
                        <option>VISHNU VIHAR-1</option>
                        <option>VISHNU VIHAR-2</option>
                        <option>VISHNU VIHAR-3</option>
                        <option>WAHADUR GARH</option>
                        <option>WAHADUR KS MAZJHAR</option>
                        <option>WAHADURPUR</option>
                        <option>WAJIDPUR</option>
                        <option>WAZEERNAGAR</option>
                        <option>YUSUF NAGAR</option>
                        <option>ZAGATPUR</option>
                        <option>ZAHIDPUR</option>
                        <option>ZALPUR</option>
                        <option>ZAMALGANJ</option>
                        <option>ZIARAT KHURMEWALI 2</option>
                        <option>ZIARAT KHURMEWALI-1</option>
                     </select>
                  </div>
               </div>
               <div class="form-group">
                  <label for="criteriavoters" class="col-sm-2 control-label">Criteria Voters: </label>
                  <div class="col-sm-10">
                     <label for="criteriavoters" class="col-sm-2 control-label">{{.CriteriaVotersCount}}</label>
                  </div>
               </div>               
               <div class="form-group">
                  <label for="criteriavoters" class="col-sm-2 control-label">Criteria Voters Percentage: </label>
                  <div class="col-sm-10">
                     <label for="criteriavoters" class="col-sm-2 control-label">{{.CriteriaVotersPercentage}}%</label>
                  </div>
               </div>
               <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                     <button type="submit" class="btn btn-danger">Calculate</button>
                     <a class="btn btn-danger" href="/election/home" style="padding-left: 11px;">Cancel</a>
                  </div>
               </div>
            </form>
         </div>
      </div>
      <!--/.col (right) -->
   </div>
   <div class="row">
      <table class="table">
         <thead>
            <tr>
               <th>Name</th>
               <th>Relation Name</th>
               <th>Gender</th>
               <th>Age</th>
               <th>Religion</th>
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
               <td>{{$record.Age}}</td>
               <td>{{$record.Religion_english}}</td>
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