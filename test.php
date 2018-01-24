<?php
/*
    
    Simple testing script

*/

define( 'SERVER', 'localhost' );
define( 'PORT', 20000 );
define( 'THUMB_WIDTH', 200 );

function println( $msg ) {
    print( $msg . "\n" );
}



$tst_path = getcwd() . '/../winsock_test/test_img/DJI_0510.JPG';

if( ($sock = socket_create( AF_INET, SOCK_DGRAM, 0) ) ){

    // headers as follows: UINT32 UINT32 CHAR[] (width, height, path)
    $msg = pack( 'NN', THUMB_WIDTH, THUMB_WIDTH * .75 ) . $tst_path;
    // $msg ='';

    println( sprintf('Request values %d, %d, %s', THUMB_WIDTH, THUMB_WIDTH * .75, $tst_path ) );
    
    $reqmsg = 'opening server request ... ';
    $res = socket_sendto($sock, $msg, strlen($msg), 0, SERVER, PORT);
    $reqmsg .= ($res) ? 'OK' : 'error';
    println( $reqmsg );

    $response;
    println( 'listening server response ... ' );
    socket_recv( $sock, $response, 2048, MSG_WAITALL );
    if( $response == '-1' ){
        println( 'got -1 response' );
    }
    else {
        println( $response );    
    }
    
}


?>