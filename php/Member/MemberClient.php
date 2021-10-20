<?php
/**
 * @date 2021/10/19
 * @author Robin
 */

namespace Member;


class MemberClient extends \Grpc\BaseStub
{
    public function __construct($hostname, $opts, $channel = null)
    {
        parent::__construct($hostname, $opts, $channel);
    }

    public function getUserInfo(GetUserInfoRequest $argument, $metadata = [], $options = [])
    {
        return $this->_simpleRequest(
            '/member.MemberService/GetUserInfo',
            $argument,
            ['Member\GetUserInfoResponse', 'decode'],
            $metadata,
            $options
        );
    }
}
