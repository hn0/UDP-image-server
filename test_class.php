<?php

require('imgsrv_client.php');

$tst_width  = 400;
$tst_height = $tst_width * .75;
$tst_path   = '/test_img/DJI_0510.JPG';

$netimage = new NetImage();

// public function GetImage( $path, $w, $h )
if( $netimage->GetImage( $tst_path, $tst_width, $tst_height ) ){
    print( "<img src=\"data:image/jpeg;base64, {$netimage->Encode64()}\" />" );
}
else {
    print( '<p>Error, could not get image from the server!</p>' );
}