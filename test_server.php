<?php
/*
    
    Simple testing script

*/

require( 'imgsrv_client.php' );
define( 'SERVER', 'localhost' );
define( 'PORT', 20000 );
define( 'THUMB_WIDTH', 800 );

function println( $msg ) {
    print( $msg . "\n" );
}

$tst_path = 'img/DJI_0510.JPG';
$out_path = 'test.jpg';

if( defined('STDIN') && sizeof($argv) > 1 ){
    $tst_path = $argv[1];    
}

if( defined('STDIN') && sizeof($argv) > 2 ){
    $out_path = $argv[2];    
}

$trim = [200,200,750,750];
$contrast = 1;

$netimage = new NetImage( SERVER, PORT, println );
$start = microtime(true);
if( $netimage->GetImage( $tst_path, THUMB_WIDTH, THUMB_WIDTH * .75, $trim, $contrast ) ){
    file_put_contents('test_results/' . $out_path, $netimage->Blob() );
}

println( sprintf( "Request time: %f", (microtime(true) - $start) ) );

?>