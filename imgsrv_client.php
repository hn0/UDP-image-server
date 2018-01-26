<?php
/*
    
    Client side class, packs request for the server and, currently provides interface for returning base64 encoded representation of the same

*/

class NetImage 
{

    public static $IMG_BUFF = 2048;
    private $server;
    private $port;

    private $image;

    public function __construct( $server='localhost', $port=20000 )
    {
        $this->server = $server;
        $this->port   = $port;

    }

    public function GetImage( $path, $w, $h )
    {
        if( ($sock = socket_create( AF_INET, SOCK_DGRAM, 0)) ) {
            $msg = pack( 'NN', $w, $h ) . $path;
            if( socket_sendto( $sock, $msg, strlen($msg), 0, $this->server, $this->port ) ){
                $bytes;
                while( socket_recv( $sock, $bytes, self::$IMG_BUFF, MSG_WAITALL ) == self::$IMG_BUFF ){
                    $this->image .= $bytes;
                }
            }
        }
        return isset( $this->image ) && $this->image != "-1";
    }

    public function Encode64()
    {
        if( isset($this->image) ){
            return base64_encode( $this->image );
        }
        return '';
    }

}