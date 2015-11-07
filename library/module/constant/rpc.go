package constant

//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
//- 各个服务ID，客户端面板ID起始10000
//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
const SERVER_GAME		=	1;	// 游戏逻辑服务器
const SERVER_CHAT		=	2; 	// 聊天服务器
const SERVER_CROSSGAME	=	100; // 跨服服务器
const SERVER_CLIENT		=	10000; // 客户端默认 refer 值
//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
//- Refer 来路ID。1～10000 客户端使用
//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
const SERVER_REFER_PLAYER	=	1;	// 客户端
//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
//- Refer 来路ID. 10001 ～ max 服务端自留地 
//-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
const SERVER_REFER_GAME	=	10001;	// 游戏逻辑服务器
const SERVER_REFER_CHAT	=	10002;	// 聊天服务器
const SERVER_REFER_CROSSGAME	=	10100;	// 跨服服务器