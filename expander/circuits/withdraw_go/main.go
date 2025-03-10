package main

import (
	"fmt"
	"math/big"
	"os"
	"reflect"

	"github.com/PolyhedraZK/ExpanderCompilerCollection/ecgo"
	"github.com/PolyhedraZK/ExpanderCompilerCollection/ecgo/integration"
	"github.com/PolyhedraZK/ExpanderCompilerCollection/test"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

func component_0_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t18, _ := big.NewInt(0).SetString("6745197990210204598374042828761989596302876299545964402857411729872131034734", 10)
	t19, _ := big.NewInt(0).SetString("426281677759936592021316809065178817848084678679510574715894138690250139748", 10)
	t20, _ := big.NewInt(0).SetString("4014188762916583598888942667424965430287497824629657219807941460227372577781", 10)
	t344 := inputs[0]
	t345 := api.Add(t344, t18)
	t346 := inputs[1]
	t347 := api.Add(t346, t19)
	t348 := inputs[2]
	t349 := api.Add(t348, t20)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t345
	outputs[1] = t347
	outputs[2] = t349
	return outputs
}
func component_1_Sigma__(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t354 := inputs[0]
	t355 := api.Mul(t354, t354)
	t356 := api.Mul(t355, t355)
	t357 := api.Mul(t356, t354)
	outputs := make([]frontend.Variable, 1)
	outputs[0] = t357
	return outputs
}
func component_2_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t21, _ := big.NewInt(0).SetString("3755116341545840759015036961635468144365099804379460727348866960676715430295", 10)
	t22, _ := big.NewInt(0).SetString("20392683181271908962657137166167696619865229065446607574667232999928814731550", 10)
	t23, _ := big.NewInt(0).SetString("6703994282500560979989445930081874901355102371090652156329919603050069367661", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t361 := api.Add(t344, t21)
	t362 := api.Add(t346, t22)
	t363 := api.Add(t348, t23)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t361
	outputs[1] = t362
	outputs[2] = t363
	return outputs
}
func component_3_Mix(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t329, _ := big.NewInt(0).SetString("18732019378264290557468133440468564866454307626475683536618613112504878618481", 10)
	t330, _ := big.NewInt(0).SetString("9131299761947733513298312097611845208338517739621853568979632113419485819303", 10)
	t331, _ := big.NewInt(0).SetString("10370080108974718697676803824769673834027675643658433702224577712625900127200", 10)
	t332, _ := big.NewInt(0).SetString("20870176810702568768751421378473869562658540583882454726129544628203806653987", 10)
	t333, _ := big.NewInt(0).SetString("10595341252162738537912664445405114076324478519622938027420701542910180337937", 10)
	t334, _ := big.NewInt(0).SetString("19705173408229649878903981084052839426532978878058043055305024233888854471533", 10)
	t335, _ := big.NewInt(0).SetString("7266061498423634438633389053804536045105766754026813321943009179476902321146", 10)
	t336, _ := big.NewInt(0).SetString("11597556804922396090267472882856054602429588299176362916247939723151043581408", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t369 := api.Mul(t331, t346)
	t370 := api.Add(t368, t369)
	t371 := api.Mul(t334, t348)
	t372 := api.Add(t370, t371)
	t373 := api.Mul(t329, t344)
	t374 := api.Add(t4, t373)
	t375 := api.Mul(t332, t346)
	t376 := api.Add(t374, t375)
	t377 := api.Mul(t335, t348)
	t378 := api.Add(t376, t377)
	t379 := api.Mul(t330, t344)
	t380 := api.Add(t4, t379)
	t381 := api.Mul(t333, t346)
	t382 := api.Add(t380, t381)
	t383 := api.Mul(t336, t348)
	t384 := api.Add(t382, t383)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t372
	outputs[1] = t378
	outputs[2] = t384
	return outputs
}
func component_4_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t24, _ := big.NewInt(0).SetString("17189230569231604821073310501737896533088589624978650476197226450738944009738", 10)
	t25, _ := big.NewInt(0).SetString("18531998296162357308313149608963848512728570123579345240911571045895174353605", 10)
	t26, _ := big.NewInt(0).SetString("4433884058681415052165697534405705901078937172224017064607454469338590163489", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t391 := api.Add(t344, t24)
	t392 := api.Add(t346, t25)
	t393 := api.Add(t348, t26)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t391
	outputs[1] = t392
	outputs[2] = t393
	return outputs
}
func component_5_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t27, _ := big.NewInt(0).SetString("8020484089444009184801117822789130075555480739986478064377452360454228170229", 10)
	t28, _ := big.NewInt(0).SetString("20560640391555251236826668015235029471365697963893708697460632109250285318704", 10)
	t29, _ := big.NewInt(0).SetString("17735423966452908760211059923359580380884879536808777323265778948947638259763", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t404 := api.Add(t344, t27)
	t406 := api.Add(t346, t28)
	t408 := api.Add(t348, t29)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t404
	outputs[1] = t406
	outputs[2] = t408
	return outputs
}
func component_6_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t30, _ := big.NewInt(0).SetString("6791331612302297428695549285132291741490338679013661880702099967749867646461", 10)
	t31, _ := big.NewInt(0).SetString("10419627351290227145210525084258167372914788967175798542355001482631316994244", 10)
	t32, _ := big.NewInt(0).SetString("6206851612052541638976352943215840028030801164970177880767418169520708772536", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t419 := api.Add(t344, t30)
	t421 := api.Add(t346, t31)
	t423 := api.Add(t348, t32)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t419
	outputs[1] = t421
	outputs[2] = t423
	return outputs
}
func component_7_Mix(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t331, _ := big.NewInt(0).SetString("10370080108974718697676803824769673834027675643658433702224577712625900127200", 10)
	t334, _ := big.NewInt(0).SetString("19705173408229649878903981084052839426532978878058043055305024233888854471533", 10)
	t337, _ := big.NewInt(0).SetString("13765730681189380936346492971955185320534160954304757809496083602133165929757", 10)
	t338, _ := big.NewInt(0).SetString("12595446607664744934103076352963528000966896978346099459720409268422440395879", 10)
	t339, _ := big.NewInt(0).SetString("20498480049173041451757161739353136932402063966867101132544382489060457121690", 10)
	t340, _ := big.NewInt(0).SetString("12226297560593729389190789373669758216633073552812492133170543943243249907657", 10)
	t341, _ := big.NewInt(0).SetString("8087150636429993556473620686397944819119746067671291185379890893406156055968", 10)
	t342, _ := big.NewInt(0).SetString("15428267695360211473228142908425586842453705255249103144570280918777118090173", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t369 := api.Mul(t331, t346)
	t370 := api.Add(t368, t369)
	t371 := api.Mul(t334, t348)
	t372 := api.Add(t370, t371)
	t427 := api.Mul(t337, t344)
	t428 := api.Add(t4, t427)
	t429 := api.Mul(t339, t346)
	t430 := api.Add(t428, t429)
	t431 := api.Mul(t341, t348)
	t432 := api.Add(t430, t431)
	t433 := api.Mul(t338, t344)
	t434 := api.Add(t4, t433)
	t435 := api.Mul(t340, t346)
	t436 := api.Add(t434, t435)
	t437 := api.Mul(t342, t348)
	t438 := api.Add(t436, t437)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t372
	outputs[1] = t432
	outputs[2] = t438
	return outputs
}
func component_8_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t103, _ := big.NewInt(0).SetString("1781874611967874592137274483616240894881315449294815307306613366069350853425", 10)
	t104, _ := big.NewInt(0).SetString("9676220459425127104563807626505378474104527268335041816433595157913150665495", 10)
	t105, _ := big.NewInt(0).SetString("8364259238812534287689210722577399963878179320345509803468849104367466297989", 10)
	t106, _ := big.NewInt(0).SetString("2889496767351495797946386949910896668575115361724249874917471657626490587069", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t443 := api.Mul(t103, t346)
	t444 := api.Add(t368, t443)
	t445 := api.Mul(t104, t348)
	t446 := api.Add(t444, t445)
	t447 := api.Mul(t344, t105)
	t448 := api.Add(t346, t447)
	t449 := api.Mul(t344, t106)
	t450 := api.Add(t348, t449)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t446
	outputs[1] = t448
	outputs[2] = t450
	return outputs
}
func component_9_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t107, _ := big.NewInt(0).SetString("15203863717131037243487133177680233750660694097162830026522190480319019526887", 10)
	t108, _ := big.NewInt(0).SetString("1645017323598148583308153743253948043010266295265950623794066679542803673813", 10)
	t109, _ := big.NewInt(0).SetString("14985926134451618201070782922146535777997354606230522118685156055564432923596", 10)
	t110, _ := big.NewInt(0).SetString("11497455747123870842609033487886196057746577750687517341166074505317007288078", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t457 := api.Mul(t107, t346)
	t458 := api.Add(t368, t457)
	t459 := api.Mul(t108, t348)
	t460 := api.Add(t458, t459)
	t461 := api.Mul(t344, t109)
	t462 := api.Add(t346, t461)
	t463 := api.Mul(t344, t110)
	t464 := api.Add(t348, t463)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t460
	outputs[1] = t462
	outputs[2] = t464
	return outputs
}
func component_10_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t111, _ := big.NewInt(0).SetString("18109765756899962487111075951493451762273621105151506450773344342109668201999", 10)
	t112, _ := big.NewInt(0).SetString("8034324828084400593020431506480243533881627849088152439427470035355284392177", 10)
	t113, _ := big.NewInt(0).SetString("16846229027008741913165717881259554980809057413299912150488284683744940628261", 10)
	t114, _ := big.NewInt(0).SetString("21835563963581578576271778192505404662763222948742168673583931448375408835935", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t470 := api.Mul(t111, t346)
	t471 := api.Add(t368, t470)
	t472 := api.Mul(t112, t348)
	t473 := api.Add(t471, t472)
	t474 := api.Mul(t344, t113)
	t475 := api.Add(t346, t474)
	t476 := api.Mul(t344, t114)
	t477 := api.Add(t348, t476)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t473
	outputs[1] = t475
	outputs[2] = t477
	return outputs
}
func component_11_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t115, _ := big.NewInt(0).SetString("21536618802882283440947141155118738832596020335348742727957480541943406874436", 10)
	t116, _ := big.NewInt(0).SetString("13397320511797493654805969878195367010267669507871486661614614086160548021432", 10)
	t117, _ := big.NewInt(0).SetString("8274817596976627060721446579061034932059250181790318658419016654356916553793", 10)
	t118, _ := big.NewInt(0).SetString("11559576119047297261718762577915230877068346446232753309523408281532457130418", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t484 := api.Mul(t115, t346)
	t485 := api.Add(t368, t484)
	t486 := api.Mul(t116, t348)
	t487 := api.Add(t485, t486)
	t490 := api.Mul(t344, t117)
	t491 := api.Add(t346, t490)
	t493 := api.Mul(t344, t118)
	t494 := api.Add(t348, t493)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t487
	outputs[1] = t491
	outputs[2] = t494
	return outputs
}
func component_12_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t119, _ := big.NewInt(0).SetString("21110548928163625108646189707151361569577559205105116148655680158775559847460", 10)
	t120, _ := big.NewInt(0).SetString("13965463506707211992011711863952040570118432896827711820318513847839923700006", 10)
	t121, _ := big.NewInt(0).SetString("2754464625251737051452042869297896380028509218065510607416300542624867449301", 10)
	t122, _ := big.NewInt(0).SetString("10907469474459001232698351613440362499830316226097001251678076978108377020171", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t501 := api.Mul(t119, t346)
	t502 := api.Add(t368, t501)
	t504 := api.Mul(t120, t348)
	t505 := api.Add(t502, t504)
	t507 := api.Mul(t344, t121)
	t508 := api.Add(t346, t507)
	t510 := api.Mul(t344, t122)
	t511 := api.Add(t348, t510)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t505
	outputs[1] = t508
	outputs[2] = t511
	return outputs
}
func component_13_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t123, _ := big.NewInt(0).SetString("20501774224204372540136096556482919283387738959798723353983096093423267639300", 10)
	t124, _ := big.NewInt(0).SetString("9836931077600326261954341466265192955109945505714894685102395567763076425240", 10)
	t125, _ := big.NewInt(0).SetString("19217533572284768010875577797906138766391845135377424890965521440233301772052", 10)
	t126, _ := big.NewInt(0).SetString("7005258728852995460900263537370745968630166959734206159957799221191925945602", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t518 := api.Mul(t123, t346)
	t519 := api.Add(t368, t518)
	t521 := api.Mul(t124, t348)
	t522 := api.Add(t519, t521)
	t525 := api.Mul(t344, t125)
	t526 := api.Add(t346, t525)
	t528 := api.Mul(t344, t126)
	t529 := api.Add(t348, t528)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t522
	outputs[1] = t526
	outputs[2] = t529
	return outputs
}
func component_14_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t127, _ := big.NewInt(0).SetString("6345451795676342424205730938660185178325967413255712040877211691532798689536", 10)
	t128, _ := big.NewInt(0).SetString("2780978923276769603084110452947415993768824535337654671457442495556365161036", 10)
	t129, _ := big.NewInt(0).SetString("219671864641846575934756268958949205252482364792826985138865722150409651877", 10)
	t130, _ := big.NewInt(0).SetString("2443931363154274626039717967689506791351357117257173081384847784325709078475", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t536 := api.Mul(t127, t346)
	t537 := api.Add(t368, t536)
	t539 := api.Mul(t128, t348)
	t540 := api.Add(t537, t539)
	t543 := api.Mul(t344, t129)
	t544 := api.Add(t346, t543)
	t546 := api.Mul(t344, t130)
	t547 := api.Add(t348, t546)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t540
	outputs[1] = t544
	outputs[2] = t547
	return outputs
}
func component_15_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t131, _ := big.NewInt(0).SetString("13124186496213605736903678544398349776579723065394336602175410821613905218508", 10)
	t132, _ := big.NewInt(0).SetString("5432513339728268829134323309369787365379820462455443204721589629977134312631", 10)
	t133, _ := big.NewInt(0).SetString("10745936869168790696368181125446125013764092826641393505115044228223535523023", 10)
	t134, _ := big.NewInt(0).SetString("2700209967286437008389190340075174766403488226669328017790667859130312864557", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t554 := api.Mul(t131, t346)
	t555 := api.Add(t368, t554)
	t557 := api.Mul(t132, t348)
	t558 := api.Add(t555, t557)
	t561 := api.Mul(t344, t133)
	t562 := api.Add(t346, t561)
	t564 := api.Mul(t344, t134)
	t565 := api.Add(t348, t564)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t558
	outputs[1] = t562
	outputs[2] = t565
	return outputs
}
func component_16_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t135, _ := big.NewInt(0).SetString("15772893083972477184537403920426585293594439809285129872672815610040350722871", 10)
	t136, _ := big.NewInt(0).SetString("21294428622740779056903376466216234290427165681731300802847694130469993394218", 10)
	t137, _ := big.NewInt(0).SetString("15894266239135468928185960163477926922877264274860345967753038330869627204155", 10)
	t138, _ := big.NewInt(0).SetString("1096368123578790517530711897777194394731212499866120053001617840145178088046", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t572 := api.Mul(t135, t346)
	t573 := api.Add(t368, t572)
	t575 := api.Mul(t136, t348)
	t576 := api.Add(t573, t575)
	t579 := api.Mul(t344, t137)
	t580 := api.Add(t346, t579)
	t582 := api.Mul(t344, t138)
	t583 := api.Add(t348, t582)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t576
	outputs[1] = t580
	outputs[2] = t583
	return outputs
}
func component_17_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t139, _ := big.NewInt(0).SetString("1394159664042366811003813388790050758063269308116252272062876498627195056527", 10)
	t140, _ := big.NewInt(0).SetString("11261056337190313066266746243632478642455050257003187980730240798531224877809", 10)
	t141, _ := big.NewInt(0).SetString("17305755215616267997146077497692988596800400998462752069352600363708883007839", 10)
	t142, _ := big.NewInt(0).SetString("15371909256746742985463109622300958997197963549518997301051533693886710333747", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t590 := api.Mul(t139, t346)
	t591 := api.Add(t368, t590)
	t593 := api.Mul(t140, t348)
	t594 := api.Add(t591, t593)
	t597 := api.Mul(t344, t141)
	t598 := api.Add(t346, t597)
	t600 := api.Mul(t344, t142)
	t601 := api.Add(t348, t600)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t594
	outputs[1] = t598
	outputs[2] = t601
	return outputs
}
func component_18_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t143, _ := big.NewInt(0).SetString("20448403594130444648089851873755778887290146036948090191937739293689284059473", 10)
	t144, _ := big.NewInt(0).SetString("4729734530435653548119746580911521748567799572047317151447278252902717458440", 10)
	t145, _ := big.NewInt(0).SetString("9055786267907928908044744667038735571363428775572377654006433176678216544138", 10)
	t146, _ := big.NewInt(0).SetString("9245235689750537947580373772395968915903822328347419898008094165262061513168", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t608 := api.Mul(t143, t346)
	t609 := api.Add(t368, t608)
	t611 := api.Mul(t144, t348)
	t612 := api.Add(t609, t611)
	t615 := api.Mul(t344, t145)
	t616 := api.Add(t346, t615)
	t618 := api.Mul(t344, t146)
	t619 := api.Add(t348, t618)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t612
	outputs[1] = t616
	outputs[2] = t619
	return outputs
}
func component_19_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t147, _ := big.NewInt(0).SetString("3259295965548895132416347844457131035605305127351914029013784648223586893840", 10)
	t148, _ := big.NewInt(0).SetString("8133110647024433575836378618144076616087915311423771001766168251715944436436", 10)
	t149, _ := big.NewInt(0).SetString("18008110744560769834041791617986172641037836309092881379393935691644464895108", 10)
	t150, _ := big.NewInt(0).SetString("9013781624325778780635119850834699693214454594410089381646984478492152387681", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t625 := api.Mul(t147, t346)
	t626 := api.Add(t368, t625)
	t627 := api.Mul(t148, t348)
	t628 := api.Add(t626, t627)
	t631 := api.Mul(t344, t149)
	t632 := api.Add(t346, t631)
	t633 := api.Mul(t344, t150)
	t634 := api.Add(t348, t633)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t628
	outputs[1] = t632
	outputs[2] = t634
	return outputs
}
func component_20_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t151, _ := big.NewInt(0).SetString("8639475724251693453868768913531642954729623102539857464903122082472741556796", 10)
	t152, _ := big.NewInt(0).SetString("20830477318165650288464577487190659978049487402162708436273498600859419634", 10)
	t153, _ := big.NewInt(0).SetString("13349403513519757309593948043861292012890478614413714204682445685718878345535", 10)
	t154, _ := big.NewInt(0).SetString("12328718012639542828603926948594616778151940577607872267472093244388211484665", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t641 := api.Mul(t151, t346)
	t642 := api.Add(t368, t641)
	t644 := api.Mul(t152, t348)
	t645 := api.Add(t642, t644)
	t646 := api.Mul(t344, t153)
	t647 := api.Add(t346, t646)
	t648 := api.Mul(t344, t154)
	t649 := api.Add(t348, t648)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t645
	outputs[1] = t647
	outputs[2] = t649
	return outputs
}
func component_21_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t155, _ := big.NewInt(0).SetString("2915193368065516044845133384670589952110028644251918175654110563684523822623", 10)
	t156, _ := big.NewInt(0).SetString("734569780368547903851295084790632331276116174575476972380730437666080976462", 10)
	t157, _ := big.NewInt(0).SetString("671279589493917786728461606950395733859229090661420264134519841071301262611", 10)
	t158, _ := big.NewInt(0).SetString("14678633946393860532975080521069035476080119750719889071999652281987539169763", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t655 := api.Mul(t155, t346)
	t656 := api.Add(t368, t655)
	t658 := api.Mul(t156, t348)
	t659 := api.Add(t656, t658)
	t661 := api.Mul(t344, t157)
	t662 := api.Add(t346, t661)
	t663 := api.Mul(t344, t158)
	t664 := api.Add(t348, t663)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t659
	outputs[1] = t662
	outputs[2] = t664
	return outputs
}
func component_22_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t159, _ := big.NewInt(0).SetString("1691723231954090840146258931861867912252544708433831341842516308673817885610", 10)
	t160, _ := big.NewInt(0).SetString("15574291717899911745152218359999334153551671302357403351163198662554477508279", 10)
	t161, _ := big.NewInt(0).SetString("5981433277656201872845331017220505919530200539512006725994262794217018602010", 10)
	t162, _ := big.NewInt(0).SetString("18156370456324591238469578107588309514554581437801913401654775491244030795770", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t671 := api.Mul(t159, t346)
	t672 := api.Add(t368, t671)
	t674 := api.Mul(t160, t348)
	t675 := api.Add(t672, t674)
	t678 := api.Mul(t344, t161)
	t679 := api.Add(t346, t678)
	t681 := api.Mul(t344, t162)
	t682 := api.Add(t348, t681)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t675
	outputs[1] = t679
	outputs[2] = t682
	return outputs
}
func component_23_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t163, _ := big.NewInt(0).SetString("1556309133439204006654419798348540449388501185001051750586019510457868307958", 10)
	t164, _ := big.NewInt(0).SetString("4356046460272772399467859547886701446225520814019018000924715176417367561817", 10)
	t165, _ := big.NewInt(0).SetString("15450880045468650144156961948500828099983553409239937576968037166948001455511", 10)
	t166, _ := big.NewInt(0).SetString("3569335951432407776495772012753227552443207946081123669782387270240663238980", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t689 := api.Mul(t163, t346)
	t690 := api.Add(t368, t689)
	t692 := api.Mul(t164, t348)
	t693 := api.Add(t690, t692)
	t696 := api.Mul(t344, t165)
	t697 := api.Add(t346, t696)
	t699 := api.Mul(t344, t166)
	t700 := api.Add(t348, t699)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t693
	outputs[1] = t697
	outputs[2] = t700
	return outputs
}
func component_24_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t167, _ := big.NewInt(0).SetString("20299619590358223273964702925591899099197268683684968495953258757381055203999", 10)
	t168, _ := big.NewInt(0).SetString("1737269388672443415630244155940415723987255613151927271717623952056489022942", 10)
	t169, _ := big.NewInt(0).SetString("7676370330863607260797103988986524817754264672351485136731920308227511577030", 10)
	t170, _ := big.NewInt(0).SetString("10764843120898224557535111936383223186451299651941198232539050093196747543756", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t706 := api.Mul(t167, t346)
	t707 := api.Add(t368, t706)
	t709 := api.Mul(t168, t348)
	t710 := api.Add(t707, t709)
	t713 := api.Mul(t344, t169)
	t714 := api.Add(t346, t713)
	t716 := api.Mul(t344, t170)
	t717 := api.Add(t348, t716)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t710
	outputs[1] = t714
	outputs[2] = t717
	return outputs
}
func component_25_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t171, _ := big.NewInt(0).SetString("2819356662200804458856836085264643083461835827345828419663815020125966978385", 10)
	t172, _ := big.NewInt(0).SetString("14230399494919677144321487695512822636538939956639271484923914516686249040244", 10)
	t173, _ := big.NewInt(0).SetString("6229792639229852919549182508857380693477833417363232050296992412866445633778", 10)
	t174, _ := big.NewInt(0).SetString("3106676750956526417925705057501789384016262285679193764776023640126964109042", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t724 := api.Mul(t171, t346)
	t725 := api.Add(t368, t724)
	t727 := api.Mul(t172, t348)
	t728 := api.Add(t725, t727)
	t731 := api.Mul(t344, t173)
	t732 := api.Add(t346, t731)
	t734 := api.Mul(t344, t174)
	t735 := api.Add(t348, t734)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t728
	outputs[1] = t732
	outputs[2] = t735
	return outputs
}
func component_26_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t175, _ := big.NewInt(0).SetString("19031174113953815401575291273416077779134839378929564662214633569481371994627", 10)
	t176, _ := big.NewInt(0).SetString("4938890649131231154991766222525002264167203279761035096310595945387423228795", 10)
	t177, _ := big.NewInt(0).SetString("9092947503088322001901942345058983345234772453274860663410155583684545688529", 10)
	t178, _ := big.NewInt(0).SetString("4443468689502285528589936084153593105296452987872236962264792108454557959607", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t742 := api.Mul(t175, t346)
	t743 := api.Add(t368, t742)
	t745 := api.Mul(t176, t348)
	t746 := api.Add(t743, t745)
	t749 := api.Mul(t344, t177)
	t750 := api.Add(t346, t749)
	t752 := api.Mul(t344, t178)
	t753 := api.Add(t348, t752)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t746
	outputs[1] = t750
	outputs[2] = t753
	return outputs
}
func component_27_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t179, _ := big.NewInt(0).SetString("13722785522864435678176292501919399406320755026890489431768679408994572946910", 10)
	t180, _ := big.NewInt(0).SetString("13256667663287458052646690425465025507007074499017697722372788741483765988169", 10)
	t181, _ := big.NewInt(0).SetString("3342109259843261627877766497639597960616083706719254912542704334341413113811", 10)
	t182, _ := big.NewInt(0).SetString("8377411907540655144604614191841171970491144397410270165752490408438880282950", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t760 := api.Mul(t179, t346)
	t761 := api.Add(t368, t760)
	t763 := api.Mul(t180, t348)
	t764 := api.Add(t761, t763)
	t767 := api.Mul(t344, t181)
	t768 := api.Add(t346, t767)
	t770 := api.Mul(t344, t182)
	t771 := api.Add(t348, t770)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t764
	outputs[1] = t768
	outputs[2] = t771
	return outputs
}
func component_28_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t183, _ := big.NewInt(0).SetString("21175860851919058796901112169110721691550903636481812384865553578742784165824", 10)
	t184, _ := big.NewInt(0).SetString("1758219250556332515525607381478749746944627538834804425466160661798760928660", 10)
	t185, _ := big.NewInt(0).SetString("8100116405804673915839318005809562313337323503890310411989391068380938049891", 10)
	t186, _ := big.NewInt(0).SetString("10950382949046383428868423373874360297216755027265677947152651089682316462002", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t778 := api.Mul(t183, t346)
	t779 := api.Add(t368, t778)
	t781 := api.Mul(t184, t348)
	t782 := api.Add(t779, t781)
	t785 := api.Mul(t344, t185)
	t786 := api.Add(t346, t785)
	t788 := api.Mul(t344, t186)
	t789 := api.Add(t348, t788)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t782
	outputs[1] = t786
	outputs[2] = t789
	return outputs
}
func component_29_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t187, _ := big.NewInt(0).SetString("2960277668778712586277871117504309767461547310299729646458954502866505810933", 10)
	t188, _ := big.NewInt(0).SetString("12436779988817213442780718350478562778741169493686625046971163883056781227217", 10)
	t189, _ := big.NewInt(0).SetString("18433130870381757859416696830699316172155927980655832716601174117670334361663", 10)
	t190, _ := big.NewInt(0).SetString("8929014056758944506773121953984691621375460981653721583817790162968859020827", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t796 := api.Mul(t187, t346)
	t797 := api.Add(t368, t796)
	t799 := api.Mul(t188, t348)
	t800 := api.Add(t797, t799)
	t803 := api.Mul(t344, t189)
	t804 := api.Add(t346, t803)
	t806 := api.Mul(t344, t190)
	t807 := api.Add(t348, t806)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t800
	outputs[1] = t804
	outputs[2] = t807
	return outputs
}
func component_30_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t191, _ := big.NewInt(0).SetString("21021117587745109604358066010067802867362858152931661595258839458778309017921", 10)
	t192, _ := big.NewInt(0).SetString("3687110520160985940053416129106142708996683054120258602350677914558228149704", 10)
	t193, _ := big.NewInt(0).SetString("80825880291398182792276850849647837369189970581427465051543823269639712237", 10)
	t194, _ := big.NewInt(0).SetString("15602858448994554323587941766253362391857349901811304586895693153675332257479", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t814 := api.Mul(t191, t346)
	t815 := api.Add(t368, t814)
	t817 := api.Mul(t192, t348)
	t818 := api.Add(t815, t817)
	t821 := api.Mul(t344, t193)
	t822 := api.Add(t346, t821)
	t824 := api.Mul(t344, t194)
	t825 := api.Add(t348, t824)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t818
	outputs[1] = t822
	outputs[2] = t825
	return outputs
}
func component_31_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t195, _ := big.NewInt(0).SetString("13135494086574956175617288396849614521078575779781791595261561845703124468256", 10)
	t196, _ := big.NewInt(0).SetString("15393949948260444958980146663126583924466023603235882001681196779684410878420", 10)
	t197, _ := big.NewInt(0).SetString("18384989275581989698635194175130733158283698892545299942532908080907204625644", 10)
	t198, _ := big.NewInt(0).SetString("485819771042979048690736635548322492095227593209398128669906407316732600888", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t832 := api.Mul(t195, t346)
	t833 := api.Add(t368, t832)
	t835 := api.Mul(t196, t348)
	t836 := api.Add(t833, t835)
	t839 := api.Mul(t344, t197)
	t840 := api.Add(t346, t839)
	t842 := api.Mul(t344, t198)
	t843 := api.Add(t348, t842)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t836
	outputs[1] = t840
	outputs[2] = t843
	return outputs
}
func component_32_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t199, _ := big.NewInt(0).SetString("3969961112111760614492622183501881958866859761703927612714294408063065400072", 10)
	t200, _ := big.NewInt(0).SetString("8752648669145926648227277846713521231276713532721674183702641053051161352313", 10)
	t201, _ := big.NewInt(0).SetString("7585110218885204638023993650637083463989720045086789711575843350789273631911", 10)
	t202, _ := big.NewInt(0).SetString("2494379627738416372577673662163694139249446937999082811387265339768290503797", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t850 := api.Mul(t199, t346)
	t851 := api.Add(t368, t850)
	t853 := api.Mul(t200, t348)
	t854 := api.Add(t851, t853)
	t857 := api.Mul(t344, t201)
	t858 := api.Add(t346, t857)
	t860 := api.Mul(t344, t202)
	t861 := api.Add(t348, t860)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t854
	outputs[1] = t858
	outputs[2] = t861
	return outputs
}
func component_33_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t203, _ := big.NewInt(0).SetString("20616688053782525026898984172292202648073622844719283906076705056594026518452", 10)
	t204, _ := big.NewInt(0).SetString("9900087106206622398227913281602779201149185950522515728836722160259149448172", 10)
	t205, _ := big.NewInt(0).SetString("11017903209339322884500424701067037363510354251034908831176623007763979729891", 10)
	t206, _ := big.NewInt(0).SetString("11242911200839364801115949018449987647748348820992122514426624004928045344694", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t868 := api.Mul(t203, t346)
	t869 := api.Add(t368, t868)
	t871 := api.Mul(t204, t348)
	t872 := api.Add(t869, t871)
	t875 := api.Mul(t344, t205)
	t876 := api.Add(t346, t875)
	t878 := api.Mul(t344, t206)
	t879 := api.Add(t348, t878)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t872
	outputs[1] = t876
	outputs[2] = t879
	return outputs
}
func component_34_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t207, _ := big.NewInt(0).SetString("19232429724858702744754565081221224741960943688294029401593672990665719107878", 10)
	t208, _ := big.NewInt(0).SetString("16765052252594983393669755070044308615954848363525024643880249721059862220578", 10)
	t209, _ := big.NewInt(0).SetString("6842036836789558363749002265840843768314388887366152991347087598440783984114", 10)
	t210, _ := big.NewInt(0).SetString("21393710061740643339940504965509850732741799591113979313939113730695101694096", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t886 := api.Mul(t207, t346)
	t887 := api.Add(t368, t886)
	t889 := api.Mul(t208, t348)
	t890 := api.Add(t887, t889)
	t893 := api.Mul(t344, t209)
	t894 := api.Add(t346, t893)
	t896 := api.Mul(t344, t210)
	t897 := api.Add(t348, t896)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t890
	outputs[1] = t894
	outputs[2] = t897
	return outputs
}
func component_35_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t211, _ := big.NewInt(0).SetString("9622969983019916007969470405619112229949366797764113862835459776222718281535", 10)
	t212, _ := big.NewInt(0).SetString("13767247240219074238794646743011288498093412255264931357766139021509967203039", 10)
	t213, _ := big.NewInt(0).SetString("20328692478494464365122435286989408673672104431805610695614028351842993934534", 10)
	t214, _ := big.NewInt(0).SetString("9073999256592381826494042793078479866030288210942587220949345879429845129344", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t904 := api.Mul(t211, t346)
	t905 := api.Add(t368, t904)
	t907 := api.Mul(t212, t348)
	t908 := api.Add(t905, t907)
	t911 := api.Mul(t344, t213)
	t912 := api.Add(t346, t911)
	t914 := api.Mul(t344, t214)
	t915 := api.Add(t348, t914)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t908
	outputs[1] = t912
	outputs[2] = t915
	return outputs
}
func component_36_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t215, _ := big.NewInt(0).SetString("8385133441250571023649882990135092851061706452670332562366981695578823064040", 10)
	t216, _ := big.NewInt(0).SetString("6908037916791839012443104181201551324508228729079993473762605932494330190638", 10)
	t217, _ := big.NewInt(0).SetString("7944824570503701879156726471230631291347547538049727334541219865644837323988", 10)
	t218, _ := big.NewInt(0).SetString("18800482911329847069658844436812670171974070641520523903011375486406401133846", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t922 := api.Mul(t215, t346)
	t923 := api.Add(t368, t922)
	t925 := api.Mul(t216, t348)
	t926 := api.Add(t923, t925)
	t929 := api.Mul(t344, t217)
	t930 := api.Add(t346, t929)
	t932 := api.Mul(t344, t218)
	t933 := api.Add(t348, t932)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t926
	outputs[1] = t930
	outputs[2] = t933
	return outputs
}
func component_37_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t219, _ := big.NewInt(0).SetString("2730366093593546914821994695117890569154816790844740397371897554795276235383", 10)
	t220, _ := big.NewInt(0).SetString("5675297339307536929988306800229752810880677519055155910685928984270724939639", 10)
	t221, _ := big.NewInt(0).SetString("8840975546939648540488041522549892926507078571712382410740665008159904893712", 10)
	t222, _ := big.NewInt(0).SetString("20979353866970550917873042661559159890255433653612953419331011151144149783744", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t940 := api.Mul(t219, t346)
	t941 := api.Add(t368, t940)
	t943 := api.Mul(t220, t348)
	t944 := api.Add(t941, t943)
	t947 := api.Mul(t344, t221)
	t948 := api.Add(t346, t947)
	t950 := api.Mul(t344, t222)
	t951 := api.Add(t348, t950)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t944
	outputs[1] = t948
	outputs[2] = t951
	return outputs
}
func component_38_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t223, _ := big.NewInt(0).SetString("516844421659953336774353304123555882256525184827876947252825317542649719056", 10)
	t224, _ := big.NewInt(0).SetString("551311298954341872590849377639279261005593012684858706728599073331951775432", 10)
	t225, _ := big.NewInt(0).SetString("21048129191517485874758270018130757373572343861561541709103852181146637709285", 10)
	t226, _ := big.NewInt(0).SetString("883108184400682278340850461255904007212979661827816162352333281411119132932", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t958 := api.Mul(t223, t346)
	t959 := api.Add(t368, t958)
	t961 := api.Mul(t224, t348)
	t962 := api.Add(t959, t961)
	t965 := api.Mul(t344, t225)
	t966 := api.Add(t346, t965)
	t968 := api.Mul(t344, t226)
	t969 := api.Add(t348, t968)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t962
	outputs[1] = t966
	outputs[2] = t969
	return outputs
}
func component_39_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t227, _ := big.NewInt(0).SetString("14420640332119892506393437524000256966511511660102357305862673030163266588863", 10)
	t228, _ := big.NewInt(0).SetString("6769807849276165954616728496863793269428109021002779834929547188571900768755", 10)
	t229, _ := big.NewInt(0).SetString("11299306373336024504558247995641644825418404376401286822173736758483745500585", 10)
	t230, _ := big.NewInt(0).SetString("3383499335919177296989189306855753260005794820125735943026533024070779082856", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t976 := api.Mul(t227, t346)
	t977 := api.Add(t368, t976)
	t979 := api.Mul(t228, t348)
	t980 := api.Add(t977, t979)
	t983 := api.Mul(t344, t229)
	t984 := api.Add(t346, t983)
	t986 := api.Mul(t344, t230)
	t987 := api.Add(t348, t986)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t980
	outputs[1] = t984
	outputs[2] = t987
	return outputs
}
func component_40_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t231, _ := big.NewInt(0).SetString("3433708777679466194488047633816494102612852206949168870493217054333441112985", 10)
	t232, _ := big.NewInt(0).SetString("13364335699281038824576139080495276061523646519119171104214550514343584904357", 10)
	t233, _ := big.NewInt(0).SetString("19088517692777810072139780055414076811493668977474813912864370395663606472109", 10)
	t234, _ := big.NewInt(0).SetString("17046893265171064448293585872818107620988569612784541924208567811178685573298", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t994 := api.Mul(t231, t346)
	t995 := api.Add(t368, t994)
	t997 := api.Mul(t232, t348)
	t998 := api.Add(t995, t997)
	t1001 := api.Mul(t344, t233)
	t1002 := api.Add(t346, t1001)
	t1004 := api.Mul(t344, t234)
	t1005 := api.Add(t348, t1004)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t998
	outputs[1] = t1002
	outputs[2] = t1005
	return outputs
}
func component_41_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t235, _ := big.NewInt(0).SetString("3339406933518442876411910401896457020433273656520834348101852668427397002466", 10)
	t236, _ := big.NewInt(0).SetString("6394754036751016627974453048774687667103663469778455952578525678514140357908", 10)
	t237, _ := big.NewInt(0).SetString("13348080011937103566625637585590574831645542599062267708945074519374215924576", 10)
	t238, _ := big.NewInt(0).SetString("2035451312942883968544771537469165070918629861375811750777728864744610711929", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1012 := api.Mul(t235, t346)
	t1013 := api.Add(t368, t1012)
	t1015 := api.Mul(t236, t348)
	t1016 := api.Add(t1013, t1015)
	t1019 := api.Mul(t344, t237)
	t1020 := api.Add(t346, t1019)
	t1022 := api.Mul(t344, t238)
	t1023 := api.Add(t348, t1022)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1016
	outputs[1] = t1020
	outputs[2] = t1023
	return outputs
}
func component_42_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t239, _ := big.NewInt(0).SetString("7534846726693802303568319129617958732413064154452139317544115737563440922906", 10)
	t240, _ := big.NewInt(0).SetString("5142893372197042264809108797404775402895973963341426202916561252529309911953", 10)
	t241, _ := big.NewInt(0).SetString("7387703761213293203195518374872886870044236674278580805224056813041998830918", 10)
	t242, _ := big.NewInt(0).SetString("9834981306855341246423988959170352646074821767371321543902587618825629388790", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1030 := api.Mul(t239, t346)
	t1031 := api.Add(t368, t1030)
	t1033 := api.Mul(t240, t348)
	t1034 := api.Add(t1031, t1033)
	t1037 := api.Mul(t344, t241)
	t1038 := api.Add(t346, t1037)
	t1040 := api.Mul(t344, t242)
	t1041 := api.Add(t348, t1040)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1034
	outputs[1] = t1038
	outputs[2] = t1041
	return outputs
}
func component_43_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t243, _ := big.NewInt(0).SetString("10591940164582290683765523873302053954617746134288371151158550854319230671848", 10)
	t244, _ := big.NewInt(0).SetString("19645940765685168416476108842047364297815786496263306942428428501384703436530", 10)
	t245, _ := big.NewInt(0).SetString("806317401532332279371557871696268272788644426105491726521005970610425656401", 10)
	t246, _ := big.NewInt(0).SetString("14873156151354922251283278949136754794279449340904101629102561195129848597881", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1048 := api.Mul(t243, t346)
	t1049 := api.Add(t368, t1048)
	t1051 := api.Mul(t244, t348)
	t1052 := api.Add(t1049, t1051)
	t1055 := api.Mul(t344, t245)
	t1056 := api.Add(t346, t1055)
	t1058 := api.Mul(t344, t246)
	t1059 := api.Add(t348, t1058)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1052
	outputs[1] = t1056
	outputs[2] = t1059
	return outputs
}
func component_44_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t247, _ := big.NewInt(0).SetString("14877529356535812861712404300630166048169645526789734524489710998713041156616", 10)
	t248, _ := big.NewInt(0).SetString("21101727915049995883360583090020188667871655700326983236468917802238514631527", 10)
	t249, _ := big.NewInt(0).SetString("8784561081435496519936150848470355611125213198581563342192869536231698468724", 10)
	t250, _ := big.NewInt(0).SetString("12951011119123862602637073643625306517125538175126787345374445023875682668190", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1066 := api.Mul(t247, t346)
	t1067 := api.Add(t368, t1066)
	t1069 := api.Mul(t248, t348)
	t1070 := api.Add(t1067, t1069)
	t1073 := api.Mul(t344, t249)
	t1074 := api.Add(t346, t1073)
	t1076 := api.Mul(t344, t250)
	t1077 := api.Add(t348, t1076)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1070
	outputs[1] = t1074
	outputs[2] = t1077
	return outputs
}
func component_45_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t251, _ := big.NewInt(0).SetString("4754486070458897643044014762078146540057558083321156154490263991438824591559", 10)
	t252, _ := big.NewInt(0).SetString("6698229600376653940889127765081219516223590790118662195996060465168245635029", 10)
	t253, _ := big.NewInt(0).SetString("3488212148323687832952214845303080200128370770801913448081307315149532795755", 10)
	t254, _ := big.NewInt(0).SetString("13395974002200754692425063613054297713599822621888055825281485401829047673168", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1084 := api.Mul(t251, t346)
	t1085 := api.Add(t368, t1084)
	t1087 := api.Mul(t252, t348)
	t1088 := api.Add(t1085, t1087)
	t1091 := api.Mul(t344, t253)
	t1092 := api.Add(t346, t1091)
	t1094 := api.Mul(t344, t254)
	t1095 := api.Add(t348, t1094)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1088
	outputs[1] = t1092
	outputs[2] = t1095
	return outputs
}
func component_46_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t255, _ := big.NewInt(0).SetString("21306313216752316778610596575521334059455780410245249300161336400126377013198", 10)
	t256, _ := big.NewInt(0).SetString("14440430794889894255165366081371645366323676828730327401596635433732808761635", 10)
	t257, _ := big.NewInt(0).SetString("11301736477249846070880364749238210747019850007649734004911360387721732439176", 10)
	t258, _ := big.NewInt(0).SetString("18529371950411247463536323927264771481897887775743653755596309214956011300885", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1102 := api.Mul(t255, t346)
	t1103 := api.Add(t368, t1102)
	t1105 := api.Mul(t256, t348)
	t1106 := api.Add(t1103, t1105)
	t1109 := api.Mul(t344, t257)
	t1110 := api.Add(t346, t1109)
	t1112 := api.Mul(t344, t258)
	t1113 := api.Add(t348, t1112)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1106
	outputs[1] = t1110
	outputs[2] = t1113
	return outputs
}
func component_47_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t259, _ := big.NewInt(0).SetString("2024094455599253391879172765188241728909648958146830531168621392830348748452", 10)
	t260, _ := big.NewInt(0).SetString("12380443335956575796199242302050308002170284713778975658193413541837749582704", 10)
	t261, _ := big.NewInt(0).SetString("17800128209140157388583882622714179816536883599865901438503119252725091065454", 10)
	t262, _ := big.NewInt(0).SetString("21045861938698937974912479796474383908520405721888783097215705657386912086696", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1120 := api.Mul(t259, t346)
	t1121 := api.Add(t368, t1120)
	t1123 := api.Mul(t260, t348)
	t1124 := api.Add(t1121, t1123)
	t1127 := api.Mul(t344, t261)
	t1128 := api.Add(t346, t1127)
	t1130 := api.Mul(t344, t262)
	t1131 := api.Add(t348, t1130)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1124
	outputs[1] = t1128
	outputs[2] = t1131
	return outputs
}
func component_48_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t263, _ := big.NewInt(0).SetString("4141409637360999331951189783363878171311106492172769273638619574221156829121", 10)
	t264, _ := big.NewInt(0).SetString("14259414300388792410641104009760954363156850399537170069218165074426770063617", 10)
	t265, _ := big.NewInt(0).SetString("4451799750330945793479450341858976120375530940735690476632525521874862862324", 10)
	t266, _ := big.NewInt(0).SetString("18172943363350781888342804719974357493732050248863214305201835660468795448831", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1138 := api.Mul(t263, t346)
	t1139 := api.Add(t368, t1138)
	t1141 := api.Mul(t264, t348)
	t1142 := api.Add(t1139, t1141)
	t1145 := api.Mul(t344, t265)
	t1146 := api.Add(t346, t1145)
	t1148 := api.Mul(t344, t266)
	t1149 := api.Add(t348, t1148)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1142
	outputs[1] = t1146
	outputs[2] = t1149
	return outputs
}
func component_49_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t267, _ := big.NewInt(0).SetString("14803601458117323257887833141099311008736410980719735518416107862729259860503", 10)
	t268, _ := big.NewInt(0).SetString("8012097819445489095043609535945175643371775681362129577114806789033825080174", 10)
	t269, _ := big.NewInt(0).SetString("20987299682170427723890380587526212844337242486458048148468388739903558239166", 10)
	t270, _ := big.NewInt(0).SetString("10548394851179037704178101661877192514367125574136880556232929084397088507285", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1156 := api.Mul(t267, t346)
	t1157 := api.Add(t368, t1156)
	t1159 := api.Mul(t268, t348)
	t1160 := api.Add(t1157, t1159)
	t1163 := api.Mul(t344, t269)
	t1164 := api.Add(t346, t1163)
	t1166 := api.Mul(t344, t270)
	t1167 := api.Add(t348, t1166)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1160
	outputs[1] = t1164
	outputs[2] = t1167
	return outputs
}
func component_50_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t271, _ := big.NewInt(0).SetString("20436799052987452454072495255981676264927711874374541657901611880206848218041", 10)
	t272, _ := big.NewInt(0).SetString("11989711640394693472854276906656379594783073287861131885588974887589308529140", 10)
	t273, _ := big.NewInt(0).SetString("18091352772795342278278111004131463236456400626592100937570367790871324385847", 10)
	t274, _ := big.NewInt(0).SetString("12711678752325475197741198013733874816358621859214685652221956581940736498324", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1174 := api.Mul(t271, t346)
	t1175 := api.Add(t368, t1174)
	t1177 := api.Mul(t272, t348)
	t1178 := api.Add(t1175, t1177)
	t1181 := api.Mul(t344, t273)
	t1182 := api.Add(t346, t1181)
	t1184 := api.Mul(t344, t274)
	t1185 := api.Add(t348, t1184)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1178
	outputs[1] = t1182
	outputs[2] = t1185
	return outputs
}
func component_51_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t275, _ := big.NewInt(0).SetString("1190440422304761108055570691102969032887211603334032397741971602684610500183", 10)
	t276, _ := big.NewInt(0).SetString("20742281673328504122132555473443044322771333000072182383854251396175500629988", 10)
	t277, _ := big.NewInt(0).SetString("6330789123996977458876730494567876598951832573056269268585355576434452265824", 10)
	t278, _ := big.NewInt(0).SetString("7613427805763613770396578102318646348515686256763144477876781927753355511242", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1192 := api.Mul(t275, t346)
	t1193 := api.Add(t368, t1192)
	t1195 := api.Mul(t276, t348)
	t1196 := api.Add(t1193, t1195)
	t1199 := api.Mul(t344, t277)
	t1200 := api.Add(t346, t1199)
	t1202 := api.Mul(t344, t278)
	t1203 := api.Add(t348, t1202)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1196
	outputs[1] = t1200
	outputs[2] = t1203
	return outputs
}
func component_52_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t279, _ := big.NewInt(0).SetString("2767787737080836074588827866493428969025899581972950836068099283611716162872", 10)
	t280, _ := big.NewInt(0).SetString("12368938928679702085904015193412499809238916971742093835750222401100611164036", 10)
	t281, _ := big.NewInt(0).SetString("2120299666226961199589805206721729429805450574305859164922602701608405684727", 10)
	t282, _ := big.NewInt(0).SetString("16101730347660865451514214922930122989814420468390642556358093789599914392935", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1210 := api.Mul(t279, t346)
	t1211 := api.Add(t368, t1210)
	t1213 := api.Mul(t280, t348)
	t1214 := api.Add(t1211, t1213)
	t1217 := api.Mul(t344, t281)
	t1218 := api.Add(t346, t1217)
	t1220 := api.Mul(t344, t282)
	t1221 := api.Add(t348, t1220)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1214
	outputs[1] = t1218
	outputs[2] = t1221
	return outputs
}
func component_53_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t283, _ := big.NewInt(0).SetString("14613859797855964370156853496634409122022020442980743716687965083719225519778", 10)
	t284, _ := big.NewInt(0).SetString("3779283189030991331381776355121793593816122884996482647339823869532343988764", 10)
	t285, _ := big.NewInt(0).SetString("16538148594031353209577287616352326794499928553504745660554665295855556894824", 10)
	t286, _ := big.NewInt(0).SetString("3123079822626887350655514696649580980677141915307255141970749507463896361323", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1228 := api.Mul(t283, t346)
	t1229 := api.Add(t368, t1228)
	t1231 := api.Mul(t284, t348)
	t1232 := api.Add(t1229, t1231)
	t1235 := api.Mul(t344, t285)
	t1236 := api.Add(t346, t1235)
	t1238 := api.Mul(t344, t286)
	t1239 := api.Add(t348, t1238)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1232
	outputs[1] = t1236
	outputs[2] = t1239
	return outputs
}
func component_54_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t287, _ := big.NewInt(0).SetString("12982425935199817815259066755446031161131158570221278702242861239646270552470", 10)
	t288, _ := big.NewInt(0).SetString("5102498747304120681063234869297561678666553390318425372362768137182642230556", 10)
	t289, _ := big.NewInt(0).SetString("5650907760235911671502574958247698947488602341810330231889326036197969521231", 10)
	t290, _ := big.NewInt(0).SetString("15311713639934636809857700294816883015313069642974788089784484331866842863071", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1246 := api.Mul(t287, t346)
	t1247 := api.Add(t368, t1246)
	t1249 := api.Mul(t288, t348)
	t1250 := api.Add(t1247, t1249)
	t1253 := api.Mul(t344, t289)
	t1254 := api.Add(t346, t1253)
	t1256 := api.Mul(t344, t290)
	t1257 := api.Add(t348, t1256)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1250
	outputs[1] = t1254
	outputs[2] = t1257
	return outputs
}
func component_55_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t291, _ := big.NewInt(0).SetString("4378917750778986566195783994933317136780665487997343184053349232575020190805", 10)
	t292, _ := big.NewInt(0).SetString("17269370569234016318347144117809553750186193189061649546246584002692850765629", 10)
	t293, _ := big.NewInt(0).SetString("15965151781956286974774343502657082669197845298829367751669865649959140668605", 10)
	t294, _ := big.NewInt(0).SetString("21450812444968239732217119395020350433942366034590850012483985750698873548994", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1264 := api.Mul(t291, t346)
	t1265 := api.Add(t368, t1264)
	t1267 := api.Mul(t292, t348)
	t1268 := api.Add(t1265, t1267)
	t1271 := api.Mul(t344, t293)
	t1272 := api.Add(t346, t1271)
	t1274 := api.Mul(t344, t294)
	t1275 := api.Add(t348, t1274)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1268
	outputs[1] = t1272
	outputs[2] = t1275
	return outputs
}
func component_56_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t295, _ := big.NewInt(0).SetString("15683936267873086453313398000666330885268595221356044868315623959998545803993", 10)
	t296, _ := big.NewInt(0).SetString("3671832753185336498356295312340707707414043518732009721061564751475499397884", 10)
	t297, _ := big.NewInt(0).SetString("8481986539959965597443698434877359782057734265717731981500359220829881743669", 10)
	t298, _ := big.NewInt(0).SetString("7660359655796884328413537474185961598411595576826789377114759090571468288601", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1282 := api.Mul(t295, t346)
	t1283 := api.Add(t368, t1282)
	t1285 := api.Mul(t296, t348)
	t1286 := api.Add(t1283, t1285)
	t1289 := api.Mul(t344, t297)
	t1290 := api.Add(t346, t1289)
	t1292 := api.Mul(t344, t298)
	t1293 := api.Add(t348, t1292)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1286
	outputs[1] = t1290
	outputs[2] = t1293
	return outputs
}
func component_57_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t299, _ := big.NewInt(0).SetString("15099124105714544055181852556690181850324058320144202151709072305108445970672", 10)
	t300, _ := big.NewInt(0).SetString("20318193804808062899310835542933059696106644785975739849404243508909313676170", 10)
	t301, _ := big.NewInt(0).SetString("19507005947491991053222274938143459936049667535869659344107661714058651936303", 10)
	t302, _ := big.NewInt(0).SetString("9680025363676779851027254588433018356491149034845693284454451321234537209837", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1300 := api.Mul(t299, t346)
	t1301 := api.Add(t368, t1300)
	t1303 := api.Mul(t300, t348)
	t1304 := api.Add(t1301, t1303)
	t1307 := api.Mul(t344, t301)
	t1308 := api.Add(t346, t1307)
	t1310 := api.Mul(t344, t302)
	t1311 := api.Add(t348, t1310)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1304
	outputs[1] = t1308
	outputs[2] = t1311
	return outputs
}
func component_58_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t303, _ := big.NewInt(0).SetString("7977470924284966780400839042253052128867651372085267651005651852743199555955", 10)
	t304, _ := big.NewInt(0).SetString("6289851497425782381089985916585292730162942529496823947960740692893599485508", 10)
	t305, _ := big.NewInt(0).SetString("1278198251448605653669861163912985025434795035476225580040678106599898395055", 10)
	t306, _ := big.NewInt(0).SetString("778822024062014472867802453882888474232798997852884487172408961114550237272", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1318 := api.Mul(t303, t346)
	t1319 := api.Add(t368, t1318)
	t1321 := api.Mul(t304, t348)
	t1322 := api.Add(t1319, t1321)
	t1325 := api.Mul(t344, t305)
	t1326 := api.Add(t346, t1325)
	t1328 := api.Mul(t344, t306)
	t1329 := api.Add(t348, t1328)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1322
	outputs[1] = t1326
	outputs[2] = t1329
	return outputs
}
func component_59_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t307, _ := big.NewInt(0).SetString("17813998309135288259967425155412879887627227853886754905994951577284709256891", 10)
	t308, _ := big.NewInt(0).SetString("13046754442426756722325203449473048800017855579216820439904651005250574252301", 10)
	t309, _ := big.NewInt(0).SetString("2675026038592592996108363640079209157158679725371291640028590665609721944662", 10)
	t310, _ := big.NewInt(0).SetString("4508630743012318612584732934628562592521561330245083297020204983532991482453", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1336 := api.Mul(t307, t346)
	t1337 := api.Add(t368, t1336)
	t1339 := api.Mul(t308, t348)
	t1340 := api.Add(t1337, t1339)
	t1343 := api.Mul(t344, t309)
	t1344 := api.Add(t346, t1343)
	t1346 := api.Mul(t344, t310)
	t1347 := api.Add(t348, t1346)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1340
	outputs[1] = t1344
	outputs[2] = t1347
	return outputs
}
func component_60_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t311, _ := big.NewInt(0).SetString("11205586019601053374384489950424904802845225981790097591516963184783396704786", 10)
	t312, _ := big.NewInt(0).SetString("3269337097979539661372044451055530562428122764943331896964292158786499210701", 10)
	t313, _ := big.NewInt(0).SetString("21019215961028087428383457025829718359262809032898137235613214997150896209535", 10)
	t314, _ := big.NewInt(0).SetString("3466829339166757648673145858981890214467602134411898125584568038757537007697", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1354 := api.Mul(t311, t346)
	t1355 := api.Add(t368, t1354)
	t1357 := api.Mul(t312, t348)
	t1358 := api.Add(t1355, t1357)
	t1361 := api.Mul(t344, t313)
	t1362 := api.Add(t346, t1361)
	t1364 := api.Mul(t344, t314)
	t1365 := api.Add(t348, t1364)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1358
	outputs[1] = t1362
	outputs[2] = t1365
	return outputs
}
func component_61_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t315, _ := big.NewInt(0).SetString("5157412242877806836300066366873354964107079264741076245467526756146318011096", 10)
	t316, _ := big.NewInt(0).SetString("21581392381591215300367149151779503009022070613614304076664343782920390616547", 10)
	t317, _ := big.NewInt(0).SetString("18549000796552159819327648418939689514195739516390499357595136551758253444650", 10)
	t318, _ := big.NewInt(0).SetString("9515161205290672029912318778766314272223114844295330905826919799686753566536", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1372 := api.Mul(t315, t346)
	t1373 := api.Add(t368, t1372)
	t1375 := api.Mul(t316, t348)
	t1376 := api.Add(t1373, t1375)
	t1379 := api.Mul(t344, t317)
	t1380 := api.Add(t346, t1379)
	t1382 := api.Mul(t344, t318)
	t1383 := api.Add(t348, t1382)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1376
	outputs[1] = t1380
	outputs[2] = t1383
	return outputs
}
func component_62_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t319, _ := big.NewInt(0).SetString("6709763924604181304099526756361626798321199970667226939575017525120090147429", 10)
	t320, _ := big.NewInt(0).SetString("3564812180471312318342772028868158337379185681492234710321340015348576731268", 10)
	t321, _ := big.NewInt(0).SetString("2715256219839290031990931607545071222786464220056110728638073108255144059506", 10)
	t322, _ := big.NewInt(0).SetString("2526648118676632885942026268297123310722360774374297527748460434510013028101", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1390 := api.Mul(t319, t346)
	t1391 := api.Add(t368, t1390)
	t1393 := api.Mul(t320, t348)
	t1394 := api.Add(t1391, t1393)
	t1397 := api.Mul(t344, t321)
	t1398 := api.Add(t346, t1397)
	t1400 := api.Mul(t344, t322)
	t1401 := api.Add(t348, t1400)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1394
	outputs[1] = t1398
	outputs[2] = t1401
	return outputs
}
func component_63_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t323, _ := big.NewInt(0).SetString("14946395762997152888563288005029334540378039755814859784393666974164235199684", 10)
	t324, _ := big.NewInt(0).SetString("8924616408420875343266627737208318913120073601143028545020037129947462534137", 10)
	t325, _ := big.NewInt(0).SetString("14553445721437460754651496265942888390087731770131124952756252097400616930608", 10)
	t326, _ := big.NewInt(0).SetString("6484523689837038546406369281981798795409487950329098695251686883211239498930", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1408 := api.Mul(t323, t346)
	t1409 := api.Add(t368, t1408)
	t1411 := api.Mul(t324, t348)
	t1412 := api.Add(t1409, t1411)
	t1415 := api.Mul(t344, t325)
	t1416 := api.Add(t346, t1415)
	t1418 := api.Mul(t344, t326)
	t1419 := api.Add(t348, t1418)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1412
	outputs[1] = t1416
	outputs[2] = t1419
	return outputs
}
func component_64_MixS(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t327, _ := big.NewInt(0).SetString("6279378546762757460220383767956301075209286500691039336178850629635359180183", 10)
	t328, _ := big.NewInt(0).SetString("3249524281869446882651222652032498789242625585725252350645660151130325444989", 10)
	t329, _ := big.NewInt(0).SetString("18732019378264290557468133440468564866454307626475683536618613112504878618481", 10)
	t330, _ := big.NewInt(0).SetString("9131299761947733513298312097611845208338517739621853568979632113419485819303", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t1426 := api.Mul(t327, t346)
	t1427 := api.Add(t368, t1426)
	t1429 := api.Mul(t328, t348)
	t1430 := api.Add(t1427, t1429)
	t1433 := api.Mul(t344, t329)
	t1434 := api.Add(t346, t1433)
	t1435 := api.Mul(t344, t330)
	t1436 := api.Add(t348, t1435)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1430
	outputs[1] = t1434
	outputs[2] = t1436
	return outputs
}
func component_65_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t90, _ := big.NewInt(0).SetString("148255380784797435050988367748108707226071678329729231552544164474530475505", 10)
	t91, _ := big.NewInt(0).SetString("12455016963320286149943199170327213031856517334199847717911791239594264576635", 10)
	t92, _ := big.NewInt(0).SetString("4938484771207094241571416021225789188526145811651959458066207028490239487168", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t1444 := api.Add(t344, t90)
	t1445 := api.Add(t346, t91)
	t1446 := api.Add(t348, t92)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1444
	outputs[1] = t1445
	outputs[2] = t1446
	return outputs
}
func component_66_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t93, _ := big.NewInt(0).SetString("10246318579378663345685131761175422014521877772325576451685137097369004581518", 10)
	t94, _ := big.NewInt(0).SetString("2049050629479134839952087472704012659976710958814656030641046436125418443803", 10)
	t95, _ := big.NewInt(0).SetString("13777389069170762688650820825296135648364766834707603999268593030539102422931", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t1456 := api.Add(t344, t93)
	t1457 := api.Add(t346, t94)
	t1458 := api.Add(t348, t95)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1456
	outputs[1] = t1457
	outputs[2] = t1458
	return outputs
}
func component_67_Ark(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t96, _ := big.NewInt(0).SetString("2293465760578772130353203454994751988060752014172004238858851708494457550991", 10)
	t97, _ := big.NewInt(0).SetString("6173354726105518526365269037588149920975300908099965898051063758804317864818", 10)
	t98, _ := big.NewInt(0).SetString("20864884888700633737572601890135683935475037549132028663329735513632822631102", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t1468 := api.Add(t344, t96)
	t1469 := api.Add(t346, t97)
	t1470 := api.Add(t348, t98)
	outputs := make([]frontend.Variable, 3)
	outputs[0] = t1468
	outputs[1] = t1469
	outputs[2] = t1470
	return outputs
}
func component_68_MixLast(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t102, _ := big.NewInt(0).SetString("7511745149465107256748700652201246547602992235352608707588321460060273774987", 10)
	t331, _ := big.NewInt(0).SetString("10370080108974718697676803824769673834027675643658433702224577712625900127200", 10)
	t334, _ := big.NewInt(0).SetString("19705173408229649878903981084052839426532978878058043055305024233888854471533", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t348 := inputs[2]
	t367 := api.Mul(t102, t344)
	t368 := api.Add(t4, t367)
	t369 := api.Mul(t331, t346)
	t370 := api.Add(t368, t369)
	t371 := api.Mul(t334, t348)
	t372 := api.Add(t370, t371)
	outputs := make([]frontend.Variable, 1)
	outputs[0] = t372
	return outputs
}
func component_69_PoseidonEx_2_1_(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t33, _ := big.NewInt(0).SetString("16375603635162350436232250364669249324451378530661474785953680978023373794530", 10)
	t34, _ := big.NewInt(0).SetString("15688345709279674878722778274755546879655509895442959219801847456408443245585", 10)
	t35, _ := big.NewInt(0).SetString("9491195295080912096808640399994744159859678118343162847585525711429214413024", 10)
	t36, _ := big.NewInt(0).SetString("9797453712978351739894993124526343599910864939600507506817907398049628087845", 10)
	t37, _ := big.NewInt(0).SetString("21481156634888978845506145026281060650315619389631972720682147891193932034748", 10)
	t38, _ := big.NewInt(0).SetString("1544695019100535789562080715491958130358622823716581449438533301216924752935", 10)
	t39, _ := big.NewInt(0).SetString("15153967549418678242792255556974876142451438236452833905885476522771426565724", 10)
	t40, _ := big.NewInt(0).SetString("4591255420184723367998678386069903388982581566230137478170120814157251999972", 10)
	t41, _ := big.NewInt(0).SetString("13993317492298544887941044850630591562583461951060762639175439957405637125554", 10)
	t42, _ := big.NewInt(0).SetString("18050986222741620548156772647408352996300510941831685700744011415483819773010", 10)
	t43, _ := big.NewInt(0).SetString("582246807524529302909723370549441534244069879807711548626660000973375204921", 10)
	t44, _ := big.NewInt(0).SetString("17980568461424306839096120761698253698461014969574413132599910426852670637994", 10)
	t45, _ := big.NewInt(0).SetString("14228661217337404173590037181281556515313880823067200751208433351015082633231", 10)
	t46, _ := big.NewInt(0).SetString("17176587110943721909591525594639263627408109053511250375171964599662347949654", 10)
	t47, _ := big.NewInt(0).SetString("7286056960291791961279922035116305681626907328744157355775762073644197019846", 10)
	t48, _ := big.NewInt(0).SetString("11801365285243706250823971466535819473941637258351304973449723129085888576630", 10)
	t49, _ := big.NewInt(0).SetString("6789889064944432682687629097717611651009674254338563170567306510098910540667", 10)
	t50, _ := big.NewInt(0).SetString("9550619200100511068539661405398488623937521959417695171688138140248257936329", 10)
	t51, _ := big.NewInt(0).SetString("16927894918204554097233146055322393983512297393314402761978026471334045088468", 10)
	t52, _ := big.NewInt(0).SetString("2296319279680349420807150717514761554038762184731526596983718190376193064033", 10)
	t53, _ := big.NewInt(0).SetString("13381111760207441008426119944140900703001726391920993676751870388659584018005", 10)
	t54, _ := big.NewInt(0).SetString("11282457978268307664923525713815776526107144144595041430117539563509678852564", 10)
	t55, _ := big.NewInt(0).SetString("17377518636062549822834113219764678554103258757534291706153558084302477704360", 10)
	t56, _ := big.NewInt(0).SetString("20529239671116714650308624442796341176059426819897849304552671207130860806391", 10)
	t57, _ := big.NewInt(0).SetString("19313513922305909359661088066839481510878680142785006144992893032981513750163", 10)
	t58, _ := big.NewInt(0).SetString("12181397983537742191390434344829585062040306747989867043080195299198026532297", 10)
	t59, _ := big.NewInt(0).SetString("11112906716400273414317383189828104351449782172976766156576450389221891985945", 10)
	t60, _ := big.NewInt(0).SetString("16412541736785056759381201344213663399564662372426071178293124552177642678859", 10)
	t61, _ := big.NewInt(0).SetString("659264346779336196861046149708262978772865549957418762539334998250261177999", 10)
	t62, _ := big.NewInt(0).SetString("4845513029979932068519665574875148103907087162327411884857282514189560116135", 10)
	t63, _ := big.NewInt(0).SetString("5002732758219210120345003630968063328669992882526477928389701063084122341769", 10)
	t64, _ := big.NewInt(0).SetString("10252016712022906174591128558929263661248150132143972390462416316600730571625", 10)
	t65, _ := big.NewInt(0).SetString("21429601688543276478479631702989513062244319445797869558505239085486171344224", 10)
	t66, _ := big.NewInt(0).SetString("11227063021005188138910539120180069062417117307677326631195927999578666832402", 10)
	t67, _ := big.NewInt(0).SetString("2254910728581601099491456127797625022511731921877856968562861178616799012230", 10)
	t68, _ := big.NewInt(0).SetString("5924174077205168234689774914167707651618793087685768535543746729243682127746", 10)
	t69, _ := big.NewInt(0).SetString("329090408153092313434075726893539446277285458579468693042578376323593473572", 10)
	t70, _ := big.NewInt(0).SetString("3484834587887234802733103827332793869706642074000786703905145704379481896136", 10)
	t71, _ := big.NewInt(0).SetString("12759747455419586364957557614124565024455324273775792120780800828643067189145", 10)
	t72, _ := big.NewInt(0).SetString("13150191605185674559081945246113753211459390086746711042772368219406961549392", 10)
	t73, _ := big.NewInt(0).SetString("6143756015450030363279441218617635078858673495963778498235578799829663351430", 10)
	t74, _ := big.NewInt(0).SetString("18969449300908196125647274430671901552593706566744295860846386166630317453793", 10)
	t75, _ := big.NewInt(0).SetString("1852637158976378935795799109534699742700007284464701345503208109137291661250", 10)
	t76, _ := big.NewInt(0).SetString("9326761420703801200266867558954051317841905707190944714132337564904087549583", 10)
	t77, _ := big.NewInt(0).SetString("6279482686602249364815416065639446422429357296367124306817890060402815786728", 10)
	t78, _ := big.NewInt(0).SetString("8520294966848398129322322020893248716223461240734329732456748763332989445897", 10)
	t79, _ := big.NewInt(0).SetString("15681345134148763222663156294793340025833734930392220652982726544070262099820", 10)
	t80, _ := big.NewInt(0).SetString("17329667728585195296928718012738338154006158317991934918090698864750378948204", 10)
	t81, _ := big.NewInt(0).SetString("13283998627857168043664255754669222819501427102611857382896531955237893912656", 10)
	t82, _ := big.NewInt(0).SetString("6734950835262505445568244961310758511728644659360842525493721393514729768139", 10)
	t83, _ := big.NewInt(0).SetString("12640921348554222969118773328433453835370715908163239963534972271298897423616", 10)
	t84, _ := big.NewInt(0).SetString("3473754313923508472440372769623619753166905053830046385167341619128450077793", 10)
	t85, _ := big.NewInt(0).SetString("15149348017909893881037206267370389784518482186719845804410708430161111942280", 10)
	t86, _ := big.NewInt(0).SetString("15095929898353593452741657787428497312742822726453112001822847009791172948206", 10)
	t87, _ := big.NewInt(0).SetString("13779749201323782722498931190091600155866019828880573899249510809182581025824", 10)
	t88, _ := big.NewInt(0).SetString("21432322857364472753097486153424499274800937939449547067783750545210710387999", 10)
	t89, _ := big.NewInt(0).SetString("16479367804307361551951437245808989924478832646635984335550324334063271392915", 10)
	t350 := inputs[2]
	t351 := inputs[0]
	t352 := inputs[1]
	sub_inputs0 := make([]frontend.Variable, 3)
	sub_inputs0[0] = t350
	sub_inputs0[1] = t351
	sub_inputs0[2] = t352
	sub_outputs0 := component_0_Ark(api, sub_inputs0)
	t358 := sub_outputs0[0]
	t359 := sub_outputs0[1]
	t360 := sub_outputs0[2]
	sub_inputs8 := make([]frontend.Variable, 1)
	sub_inputs8[0] = t358
	sub_outputs8 := component_1_Sigma__(api, sub_inputs8)
	t364 := sub_outputs8[0]
	sub_inputs9 := make([]frontend.Variable, 1)
	sub_inputs9[0] = t359
	sub_outputs9 := component_1_Sigma__(api, sub_inputs9)
	t365 := sub_outputs9[0]
	sub_inputs10 := make([]frontend.Variable, 1)
	sub_inputs10[0] = t360
	sub_outputs10 := component_1_Sigma__(api, sub_inputs10)
	t366 := sub_outputs10[0]
	sub_inputs1 := make([]frontend.Variable, 3)
	sub_inputs1[0] = t364
	sub_inputs1[1] = t365
	sub_inputs1[2] = t366
	sub_outputs1 := component_2_Ark(api, sub_inputs1)
	t385 := sub_outputs1[0]
	t386 := sub_outputs1[1]
	t387 := sub_outputs1[2]
	sub_inputs89 := make([]frontend.Variable, 3)
	sub_inputs89[0] = t385
	sub_inputs89[1] = t386
	sub_inputs89[2] = t387
	sub_outputs89 := component_3_Mix(api, sub_inputs89)
	t388 := sub_outputs89[0]
	t389 := sub_outputs89[1]
	t390 := sub_outputs89[2]
	sub_inputs11 := make([]frontend.Variable, 1)
	sub_inputs11[0] = t388
	sub_outputs11 := component_1_Sigma__(api, sub_inputs11)
	t394 := sub_outputs11[0]
	sub_inputs12 := make([]frontend.Variable, 1)
	sub_inputs12[0] = t389
	sub_outputs12 := component_1_Sigma__(api, sub_inputs12)
	t395 := sub_outputs12[0]
	sub_inputs13 := make([]frontend.Variable, 1)
	sub_inputs13[0] = t390
	sub_outputs13 := component_1_Sigma__(api, sub_inputs13)
	t396 := sub_outputs13[0]
	sub_inputs2 := make([]frontend.Variable, 3)
	sub_inputs2[0] = t394
	sub_inputs2[1] = t395
	sub_inputs2[2] = t396
	sub_outputs2 := component_4_Ark(api, sub_inputs2)
	t397 := sub_outputs2[0]
	t398 := sub_outputs2[1]
	t399 := sub_outputs2[2]
	sub_inputs90 := make([]frontend.Variable, 3)
	sub_inputs90[0] = t397
	sub_inputs90[1] = t398
	sub_inputs90[2] = t399
	sub_outputs90 := component_3_Mix(api, sub_inputs90)
	t400 := sub_outputs90[0]
	t401 := sub_outputs90[1]
	t402 := sub_outputs90[2]
	sub_inputs14 := make([]frontend.Variable, 1)
	sub_inputs14[0] = t400
	sub_outputs14 := component_1_Sigma__(api, sub_inputs14)
	t409 := sub_outputs14[0]
	sub_inputs15 := make([]frontend.Variable, 1)
	sub_inputs15[0] = t401
	sub_outputs15 := component_1_Sigma__(api, sub_inputs15)
	t410 := sub_outputs15[0]
	sub_inputs16 := make([]frontend.Variable, 1)
	sub_inputs16[0] = t402
	sub_outputs16 := component_1_Sigma__(api, sub_inputs16)
	t411 := sub_outputs16[0]
	sub_inputs3 := make([]frontend.Variable, 3)
	sub_inputs3[0] = t409
	sub_inputs3[1] = t410
	sub_inputs3[2] = t411
	sub_outputs3 := component_5_Ark(api, sub_inputs3)
	t412 := sub_outputs3[0]
	t413 := sub_outputs3[1]
	t414 := sub_outputs3[2]
	sub_inputs91 := make([]frontend.Variable, 3)
	sub_inputs91[0] = t412
	sub_inputs91[1] = t413
	sub_inputs91[2] = t414
	sub_outputs91 := component_3_Mix(api, sub_inputs91)
	t415 := sub_outputs91[0]
	t416 := sub_outputs91[1]
	t417 := sub_outputs91[2]
	sub_inputs17 := make([]frontend.Variable, 1)
	sub_inputs17[0] = t415
	sub_outputs17 := component_1_Sigma__(api, sub_inputs17)
	t424 := sub_outputs17[0]
	sub_inputs18 := make([]frontend.Variable, 1)
	sub_inputs18[0] = t416
	sub_outputs18 := component_1_Sigma__(api, sub_inputs18)
	t425 := sub_outputs18[0]
	sub_inputs19 := make([]frontend.Variable, 1)
	sub_inputs19[0] = t417
	sub_outputs19 := component_1_Sigma__(api, sub_inputs19)
	t426 := sub_outputs19[0]
	sub_inputs4 := make([]frontend.Variable, 3)
	sub_inputs4[0] = t424
	sub_inputs4[1] = t425
	sub_inputs4[2] = t426
	sub_outputs4 := component_6_Ark(api, sub_inputs4)
	t439 := sub_outputs4[0]
	t440 := sub_outputs4[1]
	t441 := sub_outputs4[2]
	sub_inputs92 := make([]frontend.Variable, 3)
	sub_inputs92[0] = t439
	sub_inputs92[1] = t440
	sub_inputs92[2] = t441
	sub_outputs92 := component_7_Mix(api, sub_inputs92)
	t442 := sub_outputs92[0]
	sub_inputs32 := make([]frontend.Variable, 1)
	sub_inputs32[0] = t442
	sub_outputs32 := component_1_Sigma__(api, sub_inputs32)
	t451 := sub_outputs32[0]
	t453 := api.Add(t451, t33)
	t454 := sub_outputs92[1]
	t455 := sub_outputs92[2]
	sub_inputs96 := make([]frontend.Variable, 3)
	sub_inputs96[0] = t453
	sub_inputs96[1] = t454
	sub_inputs96[2] = t455
	sub_outputs96 := component_8_MixS(api, sub_inputs96)
	t456 := sub_outputs96[0]
	sub_inputs33 := make([]frontend.Variable, 1)
	sub_inputs33[0] = t456
	sub_outputs33 := component_1_Sigma__(api, sub_inputs33)
	t465 := sub_outputs33[0]
	t466 := api.Add(t465, t34)
	t467 := sub_outputs96[1]
	t468 := sub_outputs96[2]
	sub_inputs97 := make([]frontend.Variable, 3)
	sub_inputs97[0] = t466
	sub_inputs97[1] = t467
	sub_inputs97[2] = t468
	sub_outputs97 := component_9_MixS(api, sub_inputs97)
	t469 := sub_outputs97[0]
	sub_inputs34 := make([]frontend.Variable, 1)
	sub_inputs34[0] = t469
	sub_outputs34 := component_1_Sigma__(api, sub_inputs34)
	t478 := sub_outputs34[0]
	t480 := api.Add(t478, t35)
	t481 := sub_outputs97[1]
	t482 := sub_outputs97[2]
	sub_inputs98 := make([]frontend.Variable, 3)
	sub_inputs98[0] = t480
	sub_inputs98[1] = t481
	sub_inputs98[2] = t482
	sub_outputs98 := component_10_MixS(api, sub_inputs98)
	t483 := sub_outputs98[0]
	sub_inputs35 := make([]frontend.Variable, 1)
	sub_inputs35[0] = t483
	sub_outputs35 := component_1_Sigma__(api, sub_inputs35)
	t495 := sub_outputs35[0]
	t496 := api.Add(t495, t36)
	t497 := sub_outputs98[1]
	t498 := sub_outputs98[2]
	sub_inputs99 := make([]frontend.Variable, 3)
	sub_inputs99[0] = t496
	sub_inputs99[1] = t497
	sub_inputs99[2] = t498
	sub_outputs99 := component_11_MixS(api, sub_inputs99)
	t499 := sub_outputs99[0]
	sub_inputs36 := make([]frontend.Variable, 1)
	sub_inputs36[0] = t499
	sub_outputs36 := component_1_Sigma__(api, sub_inputs36)
	t512 := sub_outputs36[0]
	t513 := api.Add(t512, t37)
	t514 := sub_outputs99[1]
	t515 := sub_outputs99[2]
	sub_inputs100 := make([]frontend.Variable, 3)
	sub_inputs100[0] = t513
	sub_inputs100[1] = t514
	sub_inputs100[2] = t515
	sub_outputs100 := component_12_MixS(api, sub_inputs100)
	t516 := sub_outputs100[0]
	sub_inputs37 := make([]frontend.Variable, 1)
	sub_inputs37[0] = t516
	sub_outputs37 := component_1_Sigma__(api, sub_inputs37)
	t530 := sub_outputs37[0]
	t531 := api.Add(t530, t38)
	t532 := sub_outputs100[1]
	t533 := sub_outputs100[2]
	sub_inputs101 := make([]frontend.Variable, 3)
	sub_inputs101[0] = t531
	sub_inputs101[1] = t532
	sub_inputs101[2] = t533
	sub_outputs101 := component_13_MixS(api, sub_inputs101)
	t534 := sub_outputs101[0]
	sub_inputs38 := make([]frontend.Variable, 1)
	sub_inputs38[0] = t534
	sub_outputs38 := component_1_Sigma__(api, sub_inputs38)
	t548 := sub_outputs38[0]
	t549 := api.Add(t548, t39)
	t550 := sub_outputs101[1]
	t551 := sub_outputs101[2]
	sub_inputs102 := make([]frontend.Variable, 3)
	sub_inputs102[0] = t549
	sub_inputs102[1] = t550
	sub_inputs102[2] = t551
	sub_outputs102 := component_14_MixS(api, sub_inputs102)
	t552 := sub_outputs102[0]
	sub_inputs39 := make([]frontend.Variable, 1)
	sub_inputs39[0] = t552
	sub_outputs39 := component_1_Sigma__(api, sub_inputs39)
	t566 := sub_outputs39[0]
	t567 := api.Add(t566, t40)
	t568 := sub_outputs102[1]
	t569 := sub_outputs102[2]
	sub_inputs103 := make([]frontend.Variable, 3)
	sub_inputs103[0] = t567
	sub_inputs103[1] = t568
	sub_inputs103[2] = t569
	sub_outputs103 := component_15_MixS(api, sub_inputs103)
	t570 := sub_outputs103[0]
	sub_inputs40 := make([]frontend.Variable, 1)
	sub_inputs40[0] = t570
	sub_outputs40 := component_1_Sigma__(api, sub_inputs40)
	t584 := sub_outputs40[0]
	t585 := api.Add(t584, t41)
	t586 := sub_outputs103[1]
	t587 := sub_outputs103[2]
	sub_inputs104 := make([]frontend.Variable, 3)
	sub_inputs104[0] = t585
	sub_inputs104[1] = t586
	sub_inputs104[2] = t587
	sub_outputs104 := component_16_MixS(api, sub_inputs104)
	t588 := sub_outputs104[0]
	sub_inputs41 := make([]frontend.Variable, 1)
	sub_inputs41[0] = t588
	sub_outputs41 := component_1_Sigma__(api, sub_inputs41)
	t602 := sub_outputs41[0]
	t603 := api.Add(t602, t42)
	t604 := sub_outputs104[1]
	t605 := sub_outputs104[2]
	sub_inputs105 := make([]frontend.Variable, 3)
	sub_inputs105[0] = t603
	sub_inputs105[1] = t604
	sub_inputs105[2] = t605
	sub_outputs105 := component_17_MixS(api, sub_inputs105)
	t606 := sub_outputs105[0]
	sub_inputs42 := make([]frontend.Variable, 1)
	sub_inputs42[0] = t606
	sub_outputs42 := component_1_Sigma__(api, sub_inputs42)
	t620 := sub_outputs42[0]
	t621 := api.Add(t620, t43)
	t622 := sub_outputs105[1]
	t623 := sub_outputs105[2]
	sub_inputs106 := make([]frontend.Variable, 3)
	sub_inputs106[0] = t621
	sub_inputs106[1] = t622
	sub_inputs106[2] = t623
	sub_outputs106 := component_18_MixS(api, sub_inputs106)
	t624 := sub_outputs106[0]
	sub_inputs43 := make([]frontend.Variable, 1)
	sub_inputs43[0] = t624
	sub_outputs43 := component_1_Sigma__(api, sub_inputs43)
	t635 := sub_outputs43[0]
	t636 := api.Add(t635, t44)
	t637 := sub_outputs106[1]
	t638 := sub_outputs106[2]
	sub_inputs107 := make([]frontend.Variable, 3)
	sub_inputs107[0] = t636
	sub_inputs107[1] = t637
	sub_inputs107[2] = t638
	sub_outputs107 := component_19_MixS(api, sub_inputs107)
	t639 := sub_outputs107[0]
	sub_inputs44 := make([]frontend.Variable, 1)
	sub_inputs44[0] = t639
	sub_outputs44 := component_1_Sigma__(api, sub_inputs44)
	t650 := sub_outputs44[0]
	t651 := api.Add(t650, t45)
	t652 := sub_outputs107[1]
	t653 := sub_outputs107[2]
	sub_inputs108 := make([]frontend.Variable, 3)
	sub_inputs108[0] = t651
	sub_inputs108[1] = t652
	sub_inputs108[2] = t653
	sub_outputs108 := component_20_MixS(api, sub_inputs108)
	t654 := sub_outputs108[0]
	sub_inputs45 := make([]frontend.Variable, 1)
	sub_inputs45[0] = t654
	sub_outputs45 := component_1_Sigma__(api, sub_inputs45)
	t665 := sub_outputs45[0]
	t666 := api.Add(t665, t46)
	t667 := sub_outputs108[1]
	t668 := sub_outputs108[2]
	sub_inputs109 := make([]frontend.Variable, 3)
	sub_inputs109[0] = t666
	sub_inputs109[1] = t667
	sub_inputs109[2] = t668
	sub_outputs109 := component_21_MixS(api, sub_inputs109)
	t669 := sub_outputs109[0]
	sub_inputs46 := make([]frontend.Variable, 1)
	sub_inputs46[0] = t669
	sub_outputs46 := component_1_Sigma__(api, sub_inputs46)
	t683 := sub_outputs46[0]
	t684 := api.Add(t683, t47)
	t685 := sub_outputs109[1]
	t686 := sub_outputs109[2]
	sub_inputs110 := make([]frontend.Variable, 3)
	sub_inputs110[0] = t684
	sub_inputs110[1] = t685
	sub_inputs110[2] = t686
	sub_outputs110 := component_22_MixS(api, sub_inputs110)
	t687 := sub_outputs110[0]
	sub_inputs47 := make([]frontend.Variable, 1)
	sub_inputs47[0] = t687
	sub_outputs47 := component_1_Sigma__(api, sub_inputs47)
	t701 := sub_outputs47[0]
	t702 := api.Add(t701, t48)
	t703 := sub_outputs110[1]
	t704 := sub_outputs110[2]
	sub_inputs111 := make([]frontend.Variable, 3)
	sub_inputs111[0] = t702
	sub_inputs111[1] = t703
	sub_inputs111[2] = t704
	sub_outputs111 := component_23_MixS(api, sub_inputs111)
	t705 := sub_outputs111[0]
	sub_inputs48 := make([]frontend.Variable, 1)
	sub_inputs48[0] = t705
	sub_outputs48 := component_1_Sigma__(api, sub_inputs48)
	t718 := sub_outputs48[0]
	t719 := api.Add(t718, t49)
	t720 := sub_outputs111[1]
	t721 := sub_outputs111[2]
	sub_inputs112 := make([]frontend.Variable, 3)
	sub_inputs112[0] = t719
	sub_inputs112[1] = t720
	sub_inputs112[2] = t721
	sub_outputs112 := component_24_MixS(api, sub_inputs112)
	t722 := sub_outputs112[0]
	sub_inputs49 := make([]frontend.Variable, 1)
	sub_inputs49[0] = t722
	sub_outputs49 := component_1_Sigma__(api, sub_inputs49)
	t736 := sub_outputs49[0]
	t737 := api.Add(t736, t50)
	t738 := sub_outputs112[1]
	t739 := sub_outputs112[2]
	sub_inputs113 := make([]frontend.Variable, 3)
	sub_inputs113[0] = t737
	sub_inputs113[1] = t738
	sub_inputs113[2] = t739
	sub_outputs113 := component_25_MixS(api, sub_inputs113)
	t740 := sub_outputs113[0]
	sub_inputs50 := make([]frontend.Variable, 1)
	sub_inputs50[0] = t740
	sub_outputs50 := component_1_Sigma__(api, sub_inputs50)
	t754 := sub_outputs50[0]
	t755 := api.Add(t754, t51)
	t756 := sub_outputs113[1]
	t757 := sub_outputs113[2]
	sub_inputs114 := make([]frontend.Variable, 3)
	sub_inputs114[0] = t755
	sub_inputs114[1] = t756
	sub_inputs114[2] = t757
	sub_outputs114 := component_26_MixS(api, sub_inputs114)
	t758 := sub_outputs114[0]
	sub_inputs51 := make([]frontend.Variable, 1)
	sub_inputs51[0] = t758
	sub_outputs51 := component_1_Sigma__(api, sub_inputs51)
	t772 := sub_outputs51[0]
	t773 := api.Add(t772, t52)
	t774 := sub_outputs114[1]
	t775 := sub_outputs114[2]
	sub_inputs115 := make([]frontend.Variable, 3)
	sub_inputs115[0] = t773
	sub_inputs115[1] = t774
	sub_inputs115[2] = t775
	sub_outputs115 := component_27_MixS(api, sub_inputs115)
	t776 := sub_outputs115[0]
	sub_inputs52 := make([]frontend.Variable, 1)
	sub_inputs52[0] = t776
	sub_outputs52 := component_1_Sigma__(api, sub_inputs52)
	t790 := sub_outputs52[0]
	t791 := api.Add(t790, t53)
	t792 := sub_outputs115[1]
	t793 := sub_outputs115[2]
	sub_inputs116 := make([]frontend.Variable, 3)
	sub_inputs116[0] = t791
	sub_inputs116[1] = t792
	sub_inputs116[2] = t793
	sub_outputs116 := component_28_MixS(api, sub_inputs116)
	t794 := sub_outputs116[0]
	sub_inputs53 := make([]frontend.Variable, 1)
	sub_inputs53[0] = t794
	sub_outputs53 := component_1_Sigma__(api, sub_inputs53)
	t808 := sub_outputs53[0]
	t809 := api.Add(t808, t54)
	t810 := sub_outputs116[1]
	t811 := sub_outputs116[2]
	sub_inputs117 := make([]frontend.Variable, 3)
	sub_inputs117[0] = t809
	sub_inputs117[1] = t810
	sub_inputs117[2] = t811
	sub_outputs117 := component_29_MixS(api, sub_inputs117)
	t812 := sub_outputs117[0]
	sub_inputs54 := make([]frontend.Variable, 1)
	sub_inputs54[0] = t812
	sub_outputs54 := component_1_Sigma__(api, sub_inputs54)
	t826 := sub_outputs54[0]
	t827 := api.Add(t826, t55)
	t828 := sub_outputs117[1]
	t829 := sub_outputs117[2]
	sub_inputs118 := make([]frontend.Variable, 3)
	sub_inputs118[0] = t827
	sub_inputs118[1] = t828
	sub_inputs118[2] = t829
	sub_outputs118 := component_30_MixS(api, sub_inputs118)
	t830 := sub_outputs118[0]
	sub_inputs55 := make([]frontend.Variable, 1)
	sub_inputs55[0] = t830
	sub_outputs55 := component_1_Sigma__(api, sub_inputs55)
	t844 := sub_outputs55[0]
	t845 := api.Add(t844, t56)
	t846 := sub_outputs118[1]
	t847 := sub_outputs118[2]
	sub_inputs119 := make([]frontend.Variable, 3)
	sub_inputs119[0] = t845
	sub_inputs119[1] = t846
	sub_inputs119[2] = t847
	sub_outputs119 := component_31_MixS(api, sub_inputs119)
	t848 := sub_outputs119[0]
	sub_inputs56 := make([]frontend.Variable, 1)
	sub_inputs56[0] = t848
	sub_outputs56 := component_1_Sigma__(api, sub_inputs56)
	t862 := sub_outputs56[0]
	t863 := api.Add(t862, t57)
	t864 := sub_outputs119[1]
	t865 := sub_outputs119[2]
	sub_inputs120 := make([]frontend.Variable, 3)
	sub_inputs120[0] = t863
	sub_inputs120[1] = t864
	sub_inputs120[2] = t865
	sub_outputs120 := component_32_MixS(api, sub_inputs120)
	t866 := sub_outputs120[0]
	sub_inputs57 := make([]frontend.Variable, 1)
	sub_inputs57[0] = t866
	sub_outputs57 := component_1_Sigma__(api, sub_inputs57)
	t880 := sub_outputs57[0]
	t881 := api.Add(t880, t58)
	t882 := sub_outputs120[1]
	t883 := sub_outputs120[2]
	sub_inputs121 := make([]frontend.Variable, 3)
	sub_inputs121[0] = t881
	sub_inputs121[1] = t882
	sub_inputs121[2] = t883
	sub_outputs121 := component_33_MixS(api, sub_inputs121)
	t884 := sub_outputs121[0]
	sub_inputs58 := make([]frontend.Variable, 1)
	sub_inputs58[0] = t884
	sub_outputs58 := component_1_Sigma__(api, sub_inputs58)
	t898 := sub_outputs58[0]
	t899 := api.Add(t898, t59)
	t900 := sub_outputs121[1]
	t901 := sub_outputs121[2]
	sub_inputs122 := make([]frontend.Variable, 3)
	sub_inputs122[0] = t899
	sub_inputs122[1] = t900
	sub_inputs122[2] = t901
	sub_outputs122 := component_34_MixS(api, sub_inputs122)
	t902 := sub_outputs122[0]
	sub_inputs59 := make([]frontend.Variable, 1)
	sub_inputs59[0] = t902
	sub_outputs59 := component_1_Sigma__(api, sub_inputs59)
	t916 := sub_outputs59[0]
	t917 := api.Add(t916, t60)
	t918 := sub_outputs122[1]
	t919 := sub_outputs122[2]
	sub_inputs123 := make([]frontend.Variable, 3)
	sub_inputs123[0] = t917
	sub_inputs123[1] = t918
	sub_inputs123[2] = t919
	sub_outputs123 := component_35_MixS(api, sub_inputs123)
	t920 := sub_outputs123[0]
	sub_inputs60 := make([]frontend.Variable, 1)
	sub_inputs60[0] = t920
	sub_outputs60 := component_1_Sigma__(api, sub_inputs60)
	t934 := sub_outputs60[0]
	t935 := api.Add(t934, t61)
	t936 := sub_outputs123[1]
	t937 := sub_outputs123[2]
	sub_inputs124 := make([]frontend.Variable, 3)
	sub_inputs124[0] = t935
	sub_inputs124[1] = t936
	sub_inputs124[2] = t937
	sub_outputs124 := component_36_MixS(api, sub_inputs124)
	t938 := sub_outputs124[0]
	sub_inputs61 := make([]frontend.Variable, 1)
	sub_inputs61[0] = t938
	sub_outputs61 := component_1_Sigma__(api, sub_inputs61)
	t952 := sub_outputs61[0]
	t953 := api.Add(t952, t62)
	t954 := sub_outputs124[1]
	t955 := sub_outputs124[2]
	sub_inputs125 := make([]frontend.Variable, 3)
	sub_inputs125[0] = t953
	sub_inputs125[1] = t954
	sub_inputs125[2] = t955
	sub_outputs125 := component_37_MixS(api, sub_inputs125)
	t956 := sub_outputs125[0]
	sub_inputs62 := make([]frontend.Variable, 1)
	sub_inputs62[0] = t956
	sub_outputs62 := component_1_Sigma__(api, sub_inputs62)
	t970 := sub_outputs62[0]
	t971 := api.Add(t970, t63)
	t972 := sub_outputs125[1]
	t973 := sub_outputs125[2]
	sub_inputs126 := make([]frontend.Variable, 3)
	sub_inputs126[0] = t971
	sub_inputs126[1] = t972
	sub_inputs126[2] = t973
	sub_outputs126 := component_38_MixS(api, sub_inputs126)
	t974 := sub_outputs126[0]
	sub_inputs63 := make([]frontend.Variable, 1)
	sub_inputs63[0] = t974
	sub_outputs63 := component_1_Sigma__(api, sub_inputs63)
	t988 := sub_outputs63[0]
	t989 := api.Add(t988, t64)
	t990 := sub_outputs126[1]
	t991 := sub_outputs126[2]
	sub_inputs127 := make([]frontend.Variable, 3)
	sub_inputs127[0] = t989
	sub_inputs127[1] = t990
	sub_inputs127[2] = t991
	sub_outputs127 := component_39_MixS(api, sub_inputs127)
	t992 := sub_outputs127[0]
	sub_inputs64 := make([]frontend.Variable, 1)
	sub_inputs64[0] = t992
	sub_outputs64 := component_1_Sigma__(api, sub_inputs64)
	t1006 := sub_outputs64[0]
	t1007 := api.Add(t1006, t65)
	t1008 := sub_outputs127[1]
	t1009 := sub_outputs127[2]
	sub_inputs128 := make([]frontend.Variable, 3)
	sub_inputs128[0] = t1007
	sub_inputs128[1] = t1008
	sub_inputs128[2] = t1009
	sub_outputs128 := component_40_MixS(api, sub_inputs128)
	t1010 := sub_outputs128[0]
	sub_inputs65 := make([]frontend.Variable, 1)
	sub_inputs65[0] = t1010
	sub_outputs65 := component_1_Sigma__(api, sub_inputs65)
	t1024 := sub_outputs65[0]
	t1025 := api.Add(t1024, t66)
	t1026 := sub_outputs128[1]
	t1027 := sub_outputs128[2]
	sub_inputs129 := make([]frontend.Variable, 3)
	sub_inputs129[0] = t1025
	sub_inputs129[1] = t1026
	sub_inputs129[2] = t1027
	sub_outputs129 := component_41_MixS(api, sub_inputs129)
	t1028 := sub_outputs129[0]
	sub_inputs66 := make([]frontend.Variable, 1)
	sub_inputs66[0] = t1028
	sub_outputs66 := component_1_Sigma__(api, sub_inputs66)
	t1042 := sub_outputs66[0]
	t1043 := api.Add(t1042, t67)
	t1044 := sub_outputs129[1]
	t1045 := sub_outputs129[2]
	sub_inputs130 := make([]frontend.Variable, 3)
	sub_inputs130[0] = t1043
	sub_inputs130[1] = t1044
	sub_inputs130[2] = t1045
	sub_outputs130 := component_42_MixS(api, sub_inputs130)
	t1046 := sub_outputs130[0]
	sub_inputs67 := make([]frontend.Variable, 1)
	sub_inputs67[0] = t1046
	sub_outputs67 := component_1_Sigma__(api, sub_inputs67)
	t1060 := sub_outputs67[0]
	t1061 := api.Add(t1060, t68)
	t1062 := sub_outputs130[1]
	t1063 := sub_outputs130[2]
	sub_inputs131 := make([]frontend.Variable, 3)
	sub_inputs131[0] = t1061
	sub_inputs131[1] = t1062
	sub_inputs131[2] = t1063
	sub_outputs131 := component_43_MixS(api, sub_inputs131)
	t1064 := sub_outputs131[0]
	sub_inputs68 := make([]frontend.Variable, 1)
	sub_inputs68[0] = t1064
	sub_outputs68 := component_1_Sigma__(api, sub_inputs68)
	t1078 := sub_outputs68[0]
	t1079 := api.Add(t1078, t69)
	t1080 := sub_outputs131[1]
	t1081 := sub_outputs131[2]
	sub_inputs132 := make([]frontend.Variable, 3)
	sub_inputs132[0] = t1079
	sub_inputs132[1] = t1080
	sub_inputs132[2] = t1081
	sub_outputs132 := component_44_MixS(api, sub_inputs132)
	t1082 := sub_outputs132[0]
	sub_inputs69 := make([]frontend.Variable, 1)
	sub_inputs69[0] = t1082
	sub_outputs69 := component_1_Sigma__(api, sub_inputs69)
	t1096 := sub_outputs69[0]
	t1097 := api.Add(t1096, t70)
	t1098 := sub_outputs132[1]
	t1099 := sub_outputs132[2]
	sub_inputs133 := make([]frontend.Variable, 3)
	sub_inputs133[0] = t1097
	sub_inputs133[1] = t1098
	sub_inputs133[2] = t1099
	sub_outputs133 := component_45_MixS(api, sub_inputs133)
	t1100 := sub_outputs133[0]
	sub_inputs70 := make([]frontend.Variable, 1)
	sub_inputs70[0] = t1100
	sub_outputs70 := component_1_Sigma__(api, sub_inputs70)
	t1114 := sub_outputs70[0]
	t1115 := api.Add(t1114, t71)
	t1116 := sub_outputs133[1]
	t1117 := sub_outputs133[2]
	sub_inputs134 := make([]frontend.Variable, 3)
	sub_inputs134[0] = t1115
	sub_inputs134[1] = t1116
	sub_inputs134[2] = t1117
	sub_outputs134 := component_46_MixS(api, sub_inputs134)
	t1118 := sub_outputs134[0]
	sub_inputs71 := make([]frontend.Variable, 1)
	sub_inputs71[0] = t1118
	sub_outputs71 := component_1_Sigma__(api, sub_inputs71)
	t1132 := sub_outputs71[0]
	t1133 := api.Add(t1132, t72)
	t1134 := sub_outputs134[1]
	t1135 := sub_outputs134[2]
	sub_inputs135 := make([]frontend.Variable, 3)
	sub_inputs135[0] = t1133
	sub_inputs135[1] = t1134
	sub_inputs135[2] = t1135
	sub_outputs135 := component_47_MixS(api, sub_inputs135)
	t1136 := sub_outputs135[0]
	sub_inputs72 := make([]frontend.Variable, 1)
	sub_inputs72[0] = t1136
	sub_outputs72 := component_1_Sigma__(api, sub_inputs72)
	t1150 := sub_outputs72[0]
	t1151 := api.Add(t1150, t73)
	t1152 := sub_outputs135[1]
	t1153 := sub_outputs135[2]
	sub_inputs136 := make([]frontend.Variable, 3)
	sub_inputs136[0] = t1151
	sub_inputs136[1] = t1152
	sub_inputs136[2] = t1153
	sub_outputs136 := component_48_MixS(api, sub_inputs136)
	t1154 := sub_outputs136[0]
	sub_inputs73 := make([]frontend.Variable, 1)
	sub_inputs73[0] = t1154
	sub_outputs73 := component_1_Sigma__(api, sub_inputs73)
	t1168 := sub_outputs73[0]
	t1169 := api.Add(t1168, t74)
	t1170 := sub_outputs136[1]
	t1171 := sub_outputs136[2]
	sub_inputs137 := make([]frontend.Variable, 3)
	sub_inputs137[0] = t1169
	sub_inputs137[1] = t1170
	sub_inputs137[2] = t1171
	sub_outputs137 := component_49_MixS(api, sub_inputs137)
	t1172 := sub_outputs137[0]
	sub_inputs74 := make([]frontend.Variable, 1)
	sub_inputs74[0] = t1172
	sub_outputs74 := component_1_Sigma__(api, sub_inputs74)
	t1186 := sub_outputs74[0]
	t1187 := api.Add(t1186, t75)
	t1188 := sub_outputs137[1]
	t1189 := sub_outputs137[2]
	sub_inputs138 := make([]frontend.Variable, 3)
	sub_inputs138[0] = t1187
	sub_inputs138[1] = t1188
	sub_inputs138[2] = t1189
	sub_outputs138 := component_50_MixS(api, sub_inputs138)
	t1190 := sub_outputs138[0]
	sub_inputs75 := make([]frontend.Variable, 1)
	sub_inputs75[0] = t1190
	sub_outputs75 := component_1_Sigma__(api, sub_inputs75)
	t1204 := sub_outputs75[0]
	t1205 := api.Add(t1204, t76)
	t1206 := sub_outputs138[1]
	t1207 := sub_outputs138[2]
	sub_inputs139 := make([]frontend.Variable, 3)
	sub_inputs139[0] = t1205
	sub_inputs139[1] = t1206
	sub_inputs139[2] = t1207
	sub_outputs139 := component_51_MixS(api, sub_inputs139)
	t1208 := sub_outputs139[0]
	sub_inputs76 := make([]frontend.Variable, 1)
	sub_inputs76[0] = t1208
	sub_outputs76 := component_1_Sigma__(api, sub_inputs76)
	t1222 := sub_outputs76[0]
	t1223 := api.Add(t1222, t77)
	t1224 := sub_outputs139[1]
	t1225 := sub_outputs139[2]
	sub_inputs140 := make([]frontend.Variable, 3)
	sub_inputs140[0] = t1223
	sub_inputs140[1] = t1224
	sub_inputs140[2] = t1225
	sub_outputs140 := component_52_MixS(api, sub_inputs140)
	t1226 := sub_outputs140[0]
	sub_inputs77 := make([]frontend.Variable, 1)
	sub_inputs77[0] = t1226
	sub_outputs77 := component_1_Sigma__(api, sub_inputs77)
	t1240 := sub_outputs77[0]
	t1241 := api.Add(t1240, t78)
	t1242 := sub_outputs140[1]
	t1243 := sub_outputs140[2]
	sub_inputs141 := make([]frontend.Variable, 3)
	sub_inputs141[0] = t1241
	sub_inputs141[1] = t1242
	sub_inputs141[2] = t1243
	sub_outputs141 := component_53_MixS(api, sub_inputs141)
	t1244 := sub_outputs141[0]
	sub_inputs78 := make([]frontend.Variable, 1)
	sub_inputs78[0] = t1244
	sub_outputs78 := component_1_Sigma__(api, sub_inputs78)
	t1258 := sub_outputs78[0]
	t1259 := api.Add(t1258, t79)
	t1260 := sub_outputs141[1]
	t1261 := sub_outputs141[2]
	sub_inputs142 := make([]frontend.Variable, 3)
	sub_inputs142[0] = t1259
	sub_inputs142[1] = t1260
	sub_inputs142[2] = t1261
	sub_outputs142 := component_54_MixS(api, sub_inputs142)
	t1262 := sub_outputs142[0]
	sub_inputs79 := make([]frontend.Variable, 1)
	sub_inputs79[0] = t1262
	sub_outputs79 := component_1_Sigma__(api, sub_inputs79)
	t1276 := sub_outputs79[0]
	t1277 := api.Add(t1276, t80)
	t1278 := sub_outputs142[1]
	t1279 := sub_outputs142[2]
	sub_inputs143 := make([]frontend.Variable, 3)
	sub_inputs143[0] = t1277
	sub_inputs143[1] = t1278
	sub_inputs143[2] = t1279
	sub_outputs143 := component_55_MixS(api, sub_inputs143)
	t1280 := sub_outputs143[0]
	sub_inputs80 := make([]frontend.Variable, 1)
	sub_inputs80[0] = t1280
	sub_outputs80 := component_1_Sigma__(api, sub_inputs80)
	t1294 := sub_outputs80[0]
	t1295 := api.Add(t1294, t81)
	t1296 := sub_outputs143[1]
	t1297 := sub_outputs143[2]
	sub_inputs144 := make([]frontend.Variable, 3)
	sub_inputs144[0] = t1295
	sub_inputs144[1] = t1296
	sub_inputs144[2] = t1297
	sub_outputs144 := component_56_MixS(api, sub_inputs144)
	t1298 := sub_outputs144[0]
	sub_inputs81 := make([]frontend.Variable, 1)
	sub_inputs81[0] = t1298
	sub_outputs81 := component_1_Sigma__(api, sub_inputs81)
	t1312 := sub_outputs81[0]
	t1313 := api.Add(t1312, t82)
	t1314 := sub_outputs144[1]
	t1315 := sub_outputs144[2]
	sub_inputs145 := make([]frontend.Variable, 3)
	sub_inputs145[0] = t1313
	sub_inputs145[1] = t1314
	sub_inputs145[2] = t1315
	sub_outputs145 := component_57_MixS(api, sub_inputs145)
	t1316 := sub_outputs145[0]
	sub_inputs82 := make([]frontend.Variable, 1)
	sub_inputs82[0] = t1316
	sub_outputs82 := component_1_Sigma__(api, sub_inputs82)
	t1330 := sub_outputs82[0]
	t1331 := api.Add(t1330, t83)
	t1332 := sub_outputs145[1]
	t1333 := sub_outputs145[2]
	sub_inputs146 := make([]frontend.Variable, 3)
	sub_inputs146[0] = t1331
	sub_inputs146[1] = t1332
	sub_inputs146[2] = t1333
	sub_outputs146 := component_58_MixS(api, sub_inputs146)
	t1334 := sub_outputs146[0]
	sub_inputs83 := make([]frontend.Variable, 1)
	sub_inputs83[0] = t1334
	sub_outputs83 := component_1_Sigma__(api, sub_inputs83)
	t1348 := sub_outputs83[0]
	t1349 := api.Add(t1348, t84)
	t1350 := sub_outputs146[1]
	t1351 := sub_outputs146[2]
	sub_inputs147 := make([]frontend.Variable, 3)
	sub_inputs147[0] = t1349
	sub_inputs147[1] = t1350
	sub_inputs147[2] = t1351
	sub_outputs147 := component_59_MixS(api, sub_inputs147)
	t1352 := sub_outputs147[0]
	sub_inputs84 := make([]frontend.Variable, 1)
	sub_inputs84[0] = t1352
	sub_outputs84 := component_1_Sigma__(api, sub_inputs84)
	t1366 := sub_outputs84[0]
	t1367 := api.Add(t1366, t85)
	t1368 := sub_outputs147[1]
	t1369 := sub_outputs147[2]
	sub_inputs148 := make([]frontend.Variable, 3)
	sub_inputs148[0] = t1367
	sub_inputs148[1] = t1368
	sub_inputs148[2] = t1369
	sub_outputs148 := component_60_MixS(api, sub_inputs148)
	t1370 := sub_outputs148[0]
	sub_inputs85 := make([]frontend.Variable, 1)
	sub_inputs85[0] = t1370
	sub_outputs85 := component_1_Sigma__(api, sub_inputs85)
	t1384 := sub_outputs85[0]
	t1385 := api.Add(t1384, t86)
	t1386 := sub_outputs148[1]
	t1387 := sub_outputs148[2]
	sub_inputs149 := make([]frontend.Variable, 3)
	sub_inputs149[0] = t1385
	sub_inputs149[1] = t1386
	sub_inputs149[2] = t1387
	sub_outputs149 := component_61_MixS(api, sub_inputs149)
	t1388 := sub_outputs149[0]
	sub_inputs86 := make([]frontend.Variable, 1)
	sub_inputs86[0] = t1388
	sub_outputs86 := component_1_Sigma__(api, sub_inputs86)
	t1402 := sub_outputs86[0]
	t1403 := api.Add(t1402, t87)
	t1404 := sub_outputs149[1]
	t1405 := sub_outputs149[2]
	sub_inputs150 := make([]frontend.Variable, 3)
	sub_inputs150[0] = t1403
	sub_inputs150[1] = t1404
	sub_inputs150[2] = t1405
	sub_outputs150 := component_62_MixS(api, sub_inputs150)
	t1406 := sub_outputs150[0]
	sub_inputs87 := make([]frontend.Variable, 1)
	sub_inputs87[0] = t1406
	sub_outputs87 := component_1_Sigma__(api, sub_inputs87)
	t1420 := sub_outputs87[0]
	t1421 := api.Add(t1420, t88)
	t1422 := sub_outputs150[1]
	t1423 := sub_outputs150[2]
	sub_inputs151 := make([]frontend.Variable, 3)
	sub_inputs151[0] = t1421
	sub_inputs151[1] = t1422
	sub_inputs151[2] = t1423
	sub_outputs151 := component_63_MixS(api, sub_inputs151)
	t1424 := sub_outputs151[0]
	sub_inputs88 := make([]frontend.Variable, 1)
	sub_inputs88[0] = t1424
	sub_outputs88 := component_1_Sigma__(api, sub_inputs88)
	t1437 := sub_outputs88[0]
	t1438 := api.Add(t1437, t89)
	t1439 := sub_outputs151[1]
	t1440 := sub_outputs151[2]
	sub_inputs152 := make([]frontend.Variable, 3)
	sub_inputs152[0] = t1438
	sub_inputs152[1] = t1439
	sub_inputs152[2] = t1440
	sub_outputs152 := component_64_MixS(api, sub_inputs152)
	t1441 := sub_outputs152[0]
	t1442 := sub_outputs152[1]
	t1443 := sub_outputs152[2]
	sub_inputs20 := make([]frontend.Variable, 1)
	sub_inputs20[0] = t1441
	sub_outputs20 := component_1_Sigma__(api, sub_inputs20)
	t1447 := sub_outputs20[0]
	sub_inputs21 := make([]frontend.Variable, 1)
	sub_inputs21[0] = t1442
	sub_outputs21 := component_1_Sigma__(api, sub_inputs21)
	t1448 := sub_outputs21[0]
	sub_inputs22 := make([]frontend.Variable, 1)
	sub_inputs22[0] = t1443
	sub_outputs22 := component_1_Sigma__(api, sub_inputs22)
	t1449 := sub_outputs22[0]
	sub_inputs5 := make([]frontend.Variable, 3)
	sub_inputs5[0] = t1447
	sub_inputs5[1] = t1448
	sub_inputs5[2] = t1449
	sub_outputs5 := component_65_Ark(api, sub_inputs5)
	t1450 := sub_outputs5[0]
	t1451 := sub_outputs5[1]
	t1452 := sub_outputs5[2]
	sub_inputs93 := make([]frontend.Variable, 3)
	sub_inputs93[0] = t1450
	sub_inputs93[1] = t1451
	sub_inputs93[2] = t1452
	sub_outputs93 := component_3_Mix(api, sub_inputs93)
	t1453 := sub_outputs93[0]
	t1454 := sub_outputs93[1]
	t1455 := sub_outputs93[2]
	sub_inputs23 := make([]frontend.Variable, 1)
	sub_inputs23[0] = t1453
	sub_outputs23 := component_1_Sigma__(api, sub_inputs23)
	t1459 := sub_outputs23[0]
	sub_inputs24 := make([]frontend.Variable, 1)
	sub_inputs24[0] = t1454
	sub_outputs24 := component_1_Sigma__(api, sub_inputs24)
	t1460 := sub_outputs24[0]
	sub_inputs25 := make([]frontend.Variable, 1)
	sub_inputs25[0] = t1455
	sub_outputs25 := component_1_Sigma__(api, sub_inputs25)
	t1461 := sub_outputs25[0]
	sub_inputs6 := make([]frontend.Variable, 3)
	sub_inputs6[0] = t1459
	sub_inputs6[1] = t1460
	sub_inputs6[2] = t1461
	sub_outputs6 := component_66_Ark(api, sub_inputs6)
	t1462 := sub_outputs6[0]
	t1463 := sub_outputs6[1]
	t1464 := sub_outputs6[2]
	sub_inputs94 := make([]frontend.Variable, 3)
	sub_inputs94[0] = t1462
	sub_inputs94[1] = t1463
	sub_inputs94[2] = t1464
	sub_outputs94 := component_3_Mix(api, sub_inputs94)
	t1465 := sub_outputs94[0]
	t1466 := sub_outputs94[1]
	t1467 := sub_outputs94[2]
	sub_inputs26 := make([]frontend.Variable, 1)
	sub_inputs26[0] = t1465
	sub_outputs26 := component_1_Sigma__(api, sub_inputs26)
	t1471 := sub_outputs26[0]
	sub_inputs27 := make([]frontend.Variable, 1)
	sub_inputs27[0] = t1466
	sub_outputs27 := component_1_Sigma__(api, sub_inputs27)
	t1472 := sub_outputs27[0]
	sub_inputs28 := make([]frontend.Variable, 1)
	sub_inputs28[0] = t1467
	sub_outputs28 := component_1_Sigma__(api, sub_inputs28)
	t1473 := sub_outputs28[0]
	sub_inputs7 := make([]frontend.Variable, 3)
	sub_inputs7[0] = t1471
	sub_inputs7[1] = t1472
	sub_inputs7[2] = t1473
	sub_outputs7 := component_67_Ark(api, sub_inputs7)
	t1474 := sub_outputs7[0]
	t1475 := sub_outputs7[1]
	t1476 := sub_outputs7[2]
	sub_inputs95 := make([]frontend.Variable, 3)
	sub_inputs95[0] = t1474
	sub_inputs95[1] = t1475
	sub_inputs95[2] = t1476
	sub_outputs95 := component_3_Mix(api, sub_inputs95)
	t1477 := sub_outputs95[0]
	t1478 := sub_outputs95[1]
	t1479 := sub_outputs95[2]
	sub_inputs29 := make([]frontend.Variable, 1)
	sub_inputs29[0] = t1477
	sub_outputs29 := component_1_Sigma__(api, sub_inputs29)
	t1480 := sub_outputs29[0]
	sub_inputs30 := make([]frontend.Variable, 1)
	sub_inputs30[0] = t1478
	sub_outputs30 := component_1_Sigma__(api, sub_inputs30)
	t1481 := sub_outputs30[0]
	sub_inputs31 := make([]frontend.Variable, 1)
	sub_inputs31[0] = t1479
	sub_outputs31 := component_1_Sigma__(api, sub_inputs31)
	t1482 := sub_outputs31[0]
	sub_inputs153 := make([]frontend.Variable, 3)
	sub_inputs153[0] = t1480
	sub_inputs153[1] = t1481
	sub_inputs153[2] = t1482
	sub_outputs153 := component_68_MixLast(api, sub_inputs153)
	t1483 := sub_outputs153[0]
	outputs := make([]frontend.Variable, 1)
	outputs[0] = t1483
	return outputs
}
func component_70_Poseidon_2_(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t4, _ := big.NewInt(0).SetString("0", 10)
	t351 := inputs[0]
	t352 := inputs[1]
	sub_inputs0 := make([]frontend.Variable, 3)
	sub_inputs0[0] = t351
	sub_inputs0[1] = t352
	sub_inputs0[2] = t4
	sub_outputs0 := component_69_PoseidonEx_2_1_(api, sub_inputs0)
	t1484 := sub_outputs0[0]
	outputs := make([]frontend.Variable, 1)
	outputs[0] = t1484
	return outputs
}
func component_71_DualMux__(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t2, _ := big.NewInt(0).SetString("1", 10)
	t4, _ := big.NewInt(0).SetString("0", 10)
	t344 := inputs[0]
	t346 := inputs[1]
	t1490 := inputs[2]
	t1491 := api.Sub(t2, t1490)
	t1492 := api.Mul(t1490, t1491)
	t1493 := api.Sub(t1492, t4)
	t1494 := api.Sub(t346, t344)
	t1495 := api.Mul(t1494, t1490)
	t1496 := api.Add(t1495, t344)
	t1497 := api.Sub(t344, t346)
	t1498 := api.Mul(t1497, t1490)
	t1499 := api.Add(t1498, t346)
	outputs := make([]frontend.Variable, 2)
	outputs[0] = t1496
	outputs[1] = t1499
	api.AssertIsEqual(t1493, 0)
	return outputs
}
func component_72_MerkleTreeChecker_2_(api frontend.API, leaf frontend.Variable, inputs []frontend.Variable) []frontend.Variable {
	//t1513 := leaf

	t1500 := inputs[0]
	t1501 := inputs[2]
	t1502 := inputs[4]
	sub_inputs0 := make([]frontend.Variable, 3)
	sub_inputs0[0] = t1500
	sub_inputs0[1] = t1501
	sub_inputs0[2] = t1502
	sub_outputs0 := component_71_DualMux__(api, sub_inputs0)
	t1503 := sub_outputs0[0]
	t1504 := sub_outputs0[1]
	sub_inputs2 := make([]frontend.Variable, 2)
	sub_inputs2[0] = t1503
	sub_inputs2[1] = t1504
	sub_outputs2 := component_70_Poseidon_2_(api, sub_inputs2)
	t1505 := sub_outputs2[0]
	t1506 := inputs[3]
	t1507 := inputs[5]
	sub_inputs1 := make([]frontend.Variable, 3)
	sub_inputs1[0] = t1505
	sub_inputs1[1] = t1506
	sub_inputs1[2] = t1507
	sub_outputs1 := component_71_DualMux__(api, sub_inputs1)
	t1508 := sub_outputs1[0]
	t1509 := sub_outputs1[1]
	t1510 := inputs[1]
	sub_inputs3 := make([]frontend.Variable, 2)
	sub_inputs3[0] = t1508
	sub_inputs3[1] = t1509
	sub_outputs3 := component_70_Poseidon_2_(api, sub_inputs3)
	t1511 := sub_outputs3[0]
	t1512 := api.Sub(t1510, t1511)
	outputs := make([]frontend.Variable, 0)
	api.AssertIsEqual(t1512, 0)
	return outputs
}
func component_73_Withdraw_2_(api frontend.API, inputs []frontend.Variable) []frontend.Variable {
	t1485 := inputs[2]
	sub_inputs0 := make([]frontend.Variable, 2)
	sub_inputs0[0] = t1485
	sub_inputs0[1] = t1485
	sub_outputs0 := component_70_Poseidon_2_(api, sub_inputs0)
	t1486 := sub_outputs0[0]
	t1487 := inputs[1]
	t1488 := api.Sub(t1486, t1487)
	t1489 := inputs[3]
	/*
		t1501 := inputs[4]
		t1502 := inputs[6]
		t1506 := inputs[5]
		t1507 := inputs[7]
		t1510 := inputs[0]
	*/
	sub_inputs1 := make([]frontend.Variable, 2)
	sub_inputs1[0] = t1485
	sub_inputs1[1] = t1489
	sub_outputs1 := component_70_Poseidon_2_(api, sub_inputs1)
	component_72_MerkleTreeChecker_2_(api, sub_outputs1, inputs)
	outputs := make([]frontend.Variable, 0)
	api.AssertIsEqual(t1488, 0)
	return outputs
}

type Circuit struct {
	Root          frontend.Variable `gnark:",public"`
	NullifierHash frontend.Variable `gnark:",public"`
	Nullifier     frontend.Variable
	Secret        frontend.Variable
	PathElements  [2]frontend.Variable
	PathIndices   [2]frontend.Variable
}

func (c *Circuit) Define(api frontend.API) error {
	inputs := make([]frontend.Variable, 8)
	inputs[0] = c.Root
	inputs[1] = c.NullifierHash
	inputs[2] = c.Nullifier
	inputs[3] = c.Secret
	inputs[4] = c.PathElements[0]
	inputs[5] = c.PathElements[1]
	inputs[6] = c.PathIndices[0]
	inputs[7] = c.PathIndices[1]
	component_73_Withdraw_2_(api, inputs)
	return nil
}
func main() {
	circuit, _ := ecgo.Compile(ecc.BN254.ScalarField(), &Circuit{})
	c := circuit.GetLayeredCircuit()

	// Convertir los valores hexadecimales a big.Int
	root, _ := new(big.Int).SetString("11aecbbfb437e5677960ee0a9cf0e43975214b8c0cfe0327d2f618e73c05c5ea", 16)
	nullifierHash, _ := new(big.Int).SetString("157cef5f4ba2f139aecb41b479d2efbc6d888aa89a4ff10ec6850c69829f960f", 16)
	nullifier, _ := new(big.Int).SetString("0990fb0fa550d25d6ef750aa84898d944af9568656c913d14056615d86dddc54", 16)
	secret, _ := new(big.Int).SetString("07ecd6ff9737eb11d19e77a84bf2d9269a72569fd656ee7773ea4a5f3cbc321d", 16)
	pathElement1, _ := new(big.Int).SetString("1d83cc0d4195d4cc2315f3741630e89c22600c66c88684ed2b6e579472700dbb", 16)
	pathElement2, _ := new(big.Int).SetString("22ad4ea9d906223178e5e07ce96027769ee28e66bcfc00237e69c08845cd3972", 16)

	// Asignar los valores a la estructura assignment
	assignment := &Circuit{
		Root:          root,
		NullifierHash: nullifierHash,
		Nullifier:     nullifier,
		Secret:        secret,
		PathElements:  [2]frontend.Variable{pathElement1, pathElement2},
		PathIndices:   [2]frontend.Variable{big.NewInt(1), big.NewInt(0)},
	}

	// Mostrar la estructura asignada
	fmt.Printf("Assignment: %+v\n", assignment)

	v := reflect.ValueOf(*assignment)
	fmt.Println("Número de campos en assignment:", v.NumField())
	fmt.Println("Número de variables en el circuito:", circuit.GetLayeredCircuit())
	fmt.Println("Número de restricciones en el circuito:", circuit.GetInputSolver())

	os.WriteFile("circuit.txt", c.Serialize(), 0o644)
	inputSolver := circuit.GetInputSolver()
	os.WriteFile("inputsolver.txt", inputSolver.Serialize(), 0o644)
	witness, _ := inputSolver.SolveInputAuto(assignment)
	os.WriteFile("witness.txt", witness.Serialize(), 0o644)
	if !test.CheckCircuit(c, witness) {
		panic("verification failed")
	}

	circuit_name := "mio"
	circuit_dir := "circuit.txt"
	max_concurrency := 0
	prover, err := integration.NewProver(circuit_dir, circuit_name, max_concurrency, true)
	if err != nil {
		panic(err)
	}
	witnessData, err := os.ReadFile("witness.txt")
	if err != nil {
		panic(err)
	}

	println("Generating proof...")
	proof, err := prover.Prove(witnessData)
	if err != nil {
		panic(err)
	}
	println("proof:", proof[:8])
	os.WriteFile("proof.txt", proof, 0o644)

	println("Verifying proof...")
	result, err := prover.Verify(witnessData, proof)
	if err != nil {
		panic(err)
	}
	println("verification result (expecting true):", result)

	println("Verify invalid proof...")
	invalid_proof := make([]byte, len(proof))
	copy(invalid_proof, proof)
	// flip a bit
	invalid_proof[0] ^= 1
	result, _ = prover.Verify(witnessData, invalid_proof)
	println("verification result (expecting false):", result)

}
