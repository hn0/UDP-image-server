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
    private $log_fnc;

    public function __construct( $server='localhost', $port=20000, $log_fnc=null )
    {
        $this->server  = $server;
        $this->port    = $port;
        $this->log_fnc = $log_fnc;
    }

    public function GetImage( $path, $w, $h )
    {
        if( ($sock = socket_create( AF_INET, SOCK_DGRAM, 0)) ) {
            $this->log( sprintf( 'Server msg: w:%d, h:%d, path:%s', $w, $h, $path ) );
            $msg = pack( 'NN', $w, $h ) . $path;
            if( socket_sendto( $sock, $msg, strlen($msg), 0, $this->server, $this->port ) ){
            $this->log( sprintf( 'Socket is open now. Sending msg to: to %s:%s', $this->server, $this->port ) );
                $bytes;
                while( socket_recv( $sock, $bytes, self::$IMG_BUFF, MSG_WAITALL ) == self::$IMG_BUFF ){
                    $this->image .= $bytes;
                }
                if( $this->image && $this->image != -1 ){
                    $this->log( sprintf( 'Got response of size:', strlen($this->image) ) );
                }
                else{
                    $this->log( 'Request failed' );
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

    public function Blob()
    {
        if( isset($this->image) ){
            return $this->image;
        }
        return '';
    }

    private function log($msg)
    {
        if( $this->log_fnc ){
            call_user_func( $this->log_fnc, $msg );
        }
    }

}