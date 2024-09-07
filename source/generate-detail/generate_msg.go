package generatedetail

import (
	"fmt"
	"strings"
)

// 根据 name 生成 C++ 代码
func GenerateMsgCode(name string) string {

	// 生成 s2c 消息ID
	s2cMsgID := "s2c_" + name

	// 生成 s2c 消息结构体
	s2cStructName := "STU_" + strings.ToUpper(s2cMsgID)

	// 生成 c2s 消息ID
	c2sMsgID := "c2s_" + name

	// 生成 c2s 消息结构体
	c2sStructName := "STU_" + strings.ToUpper(c2sMsgID)

	// 生成 C++ 的 .h 代码
	cppCode1 := fmt.Sprintf(`
#pragma once
// 文件说明：定义消息 %s // 1

// s2c 系列 -----------------------------------------------------------------------------------------------------------
// 自动生成的消息 ID
enum s2c_CLIENT_EXTEND_PROTOCOL
{
	%s, // 2
};

// 自动生成的结构体
struct %s : CLIENT_EXTEND_HEADER // 3
{
	uint32_t nID;             // Npc的唯一ID

	// TODO: 定义结构体的成员变量
};

// 单发消息
void KPlayer::send_%s() // 4
{
	%s sendMsg; // 5
	// todo: 填充 sendMsg 的数据

	// 发送消息
	sendMsg.SetProtocolHeader(%s, sizeof(%s) - 1); // 6 7
	g_pServer->PackDataToClient(m_nNetConnectIdx, &sendMsg, sizeof(sendMsg));
}

// 群发消息
void KNpc::broadcast_%s() // 8
{
	%s sync; // 9
	sync.SetProtocolHeader(%s, sizeof(%s) -1); // 10 11
	sync.nID = m_dwID;
	// todo: 填充 sync 的数据

	// 广播消息
	int nMaxCount = MAX_PLAYER;//MAX_BROADCAST_COUNT;
	BROADCAST_REGION(&sync, sizeof(%s), nMaxCount, m_Index); // 12
}

// 客户端处理
// 注册 s2c_extend 相关协议处理
void KProtocolProcess::registerExtendMessageHandler()
{
	m_arrExtendFunctions[s2c_CLIENT_EXTEND_PROTOCOL::%s] = &KProtocolProcess::%s; // 13 14
}

bool KProtocolProcess::%s(BYTE* pMsg) // 15
{
    if (pMsg == NULL)
        return false;

    %s* pMsgInfo = (%s*)pMsg; // 16 17

	// 处理方式1
    // int nIdx = NpcSet.SearchID(pMsgInfo->ID);
    // if (!nIdx)
    // {
    //     return false;
    // }
    // Npc[nIdx];

	// 处理方式2
    // Player[CLIENT_PLAYER_INDEX];

    return true;
}

`,
		name,          // 1
		s2cMsgID,      // 2
		s2cStructName, // 3
		s2cMsgID,      // 4
		s2cStructName, // 5
		s2cMsgID,      // 6
		s2cStructName, // 7
		s2cMsgID,      // 8
		s2cStructName, // 9
		s2cMsgID,      // 10
		s2cStructName, // 11
		s2cStructName, // 12
		s2cMsgID,      // 13
		s2cMsgID,      // 14
		s2cMsgID,      // 15
		s2cStructName, // 16
		s2cStructName) // 17

	cppCode2 := fmt.Sprintf(`
// c2s 系列 -----------------------------------------------------------------------------------------------------------
// 自动生成的消息 ID
enum c2s_CLIENT_EXTEND_PROTOCOL
{
	%s, // 1
};

// 自动生成的结构体
struct %s : CLIENT_EXTEND_HEADER // 2
{
	// TODO: 定义结构体的成员变量
};

// 客户端发送消息
{
	%s sendMsg;  // 3
	sendMsg.SetProtocolHeader(%s, sizeof(%s) - 1);  // 4 5
	// todo: 填充 sendMsg 的数据
	
	if (g_pClient)
		g_pClient->SendPackToServer((char*)&sendMsg, sizeof(sendMsg));
}

// 服务器 %s 消息处理 // 6
void KPlayer::deal_%s(%s* pMsgInfo) // 7 8
{
	if (pMsgInfo == NULL)
	{
		return;
	}

	// todo: 处理消息
}

// 服务器处理消息分发
void KProtocolProcess::c2sClientExtendProcess(int nIndex, BYTE * pMsg, int nSize)
{
	switch ( pHeader->ProtocalType )
	{
	case %s: // 9
	{
		if (pMsg == NULL)
		{
			break;
		}

		if (nIndex <= 0 || nIndex >=MAX_PLAYER)
		{
			break;
		}

		%s* pMsgInfo = (%s*)pMsg; // 10 11
		Player[nIndex].deal_%s(pMsgInfo); // 12
	}
	break;
	defualt:
		break;
	}
}
`,
		c2sMsgID,      // 1
		c2sStructName, // 2
		c2sStructName, // 3
		c2sMsgID,      // 4
		c2sStructName, // 5
		name,          // 6
		c2sMsgID,      // 7
		c2sStructName, // 8
		c2sMsgID,      // 9
		c2sStructName, // 10
		c2sStructName, // 11
		c2sMsgID)      // 12

	return cppCode1 + cppCode2
}
