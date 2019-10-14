<map version="freeplane 1.7.0">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="GoBandit&#xa;App Architecture" FOLDED="false" ID="ID_157581235" CREATED="1566373505402" MODIFIED="1566375010337" STYLE="oval">
<font SIZE="18"/>
<hook NAME="MapStyle">
    <properties edgeColorConfiguration="#808080ff,#ff0000ff,#0000ffff,#00ff00ff,#ff00ffff,#00ffffff,#7c0000ff,#00007cff,#007c00ff,#7c007cff,#007c7cff,#7c7c00ff" show_note_icons="true" fit_to_viewport="false"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node" STYLE="oval" UNIFORM_SHAPE="true" VGAP_QUANTITY="24.0 pt">
<font SIZE="24"/>
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="default" ICON_SIZE="12.0 pt" COLOR="#000000" STYLE="fork">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.attributes">
<font SIZE="9"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.note" COLOR="#000000" BACKGROUND_COLOR="#ffffff" TEXT_ALIGN="LEFT"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000" STYLE="oval" SHAPE_HORIZONTAL_MARGIN="10.0 pt" SHAPE_VERTICAL_MARGIN="10.0 pt">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font SIZE="16"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font SIZE="14"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font SIZE="12"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font SIZE="10"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,5"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,6"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,7"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,8"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,9"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,10"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,11"/>
</stylenode>
</stylenode>
</map_styles>
</hook>
<hook NAME="AutomaticEdgeColor" COUNTER="10" RULE="ON_BRANCH_CREATION"/>
<node TEXT="Game" POSITION="right" ID="ID_929973716" CREATED="1566373710281" MODIFIED="1566375013358" HGAP_QUANTITY="11.000000089406964 pt" VSHIFT_QUANTITY="-35.99999892711642 pt">
<edge COLOR="#ff0000"/>
<richcontent TYPE="NOTE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Game is the main object in the application. It intialises the SDL2 library, creates Window, Surface, and Renderer.
    </p>
  </body>
</html>
</richcontent>
<node TEXT="Board" ID="ID_1773603739" CREATED="1566373774429" MODIFIED="1566406112668" HGAP_QUANTITY="16.249999932944775 pt" VSHIFT_QUANTITY="-29.24999912828209 pt"><richcontent TYPE="NOTE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Board is the main view portal, and works with Tileset interfaces ( see Tileset )
    </p>
    <p>
      You can think of the Board as the View in the MVC (Modal, View, Controller) pattern.
    </p>
  </body>
</html>
</richcontent>
<node TEXT="Tileset" ID="ID_1501374144" CREATED="1566373791476" MODIFIED="1566406118581">
<node TEXT="Tile" ID="ID_1703456710" CREATED="1566406108452" MODIFIED="1566406124946" HGAP_QUANTITY="21.49999977648259 pt" VSHIFT_QUANTITY="-22.499999329447768 pt"/>
</node>
<node TEXT="Scorecard" ID="ID_76095766" CREATED="1566374122367" MODIFIED="1566374129157">
<node TEXT="Scorer" ID="ID_680921767" CREATED="1567794124491" MODIFIED="1567794316535"><richcontent TYPE="NOTE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Scorer is accessed via the Scorecard I/F and provides a set of methods for matching sets in the rows and columns of the board, and then evaluating their value.
    </p>
    <p>
      Scorecard must return a score value, which can then be used by the Board to update the score&#160;
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
</node>
<node TEXT="" ID="ID_1382479871" CREATED="1566406219945" MODIFIED="1566406219945"/>
<node TEXT="Controls" ID="ID_1265083027" CREATED="1566374088952" MODIFIED="1566374764202"/>
</node>
<node TEXT="Configuration" POSITION="right" ID="ID_588373447" CREATED="1566406101038" MODIFIED="1566406193011">
<edge COLOR="#007c00"/>
</node>
<node TEXT="Main loop" POSITION="right" ID="ID_1590586669" CREATED="1566406105605" MODIFIED="1566406261409">
<edge COLOR="#7c007c"/>
</node>
<node TEXT="" POSITION="left" ID="ID_1799308736" CREATED="1566406097712" MODIFIED="1566406097714">
<edge COLOR="#00007c"/>
</node>
</node>
</map>
