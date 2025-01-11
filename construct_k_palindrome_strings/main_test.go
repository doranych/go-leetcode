package construct_k_palindrome_strings

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"annabelle", 2}, true},
		{"", args{"leetcode", 3}, false},
		{"", args{"true", 4}, true},
		{"", args{"ab", 10}, false},
		{"", args{"a", 1}, true},
		{"", args{"yiipem", 3}, false},
		{"", args{"abcdefghijklmnopqrstuvwxyz", 25}, false},
		{"", args{"aabcdefghijklmnopqrstuvwxy", 24}, true},
		{"", args{"rraceca", 3}, true},
		{"", args{"dwhjvyplfsdbgjxqcfvngjsowlmjpkmuasvfvbvvrmnectsrbyfomxskpajfowsnfpkxdmqcifozydqqzeflvdulslyxpyejxjvpnrkqnlrchpzaaaqsxabpgtrssnvtcathbbywnzuszajwpvpqavewpqjxzltijwnpjdhhmezcfaazghlcpylwyylpkncoqlseqizirlxuzlicmezlrbptgmcixhruqtrbtwuebsstsmsttzexdmpkfuukevyepkzmfuetklhbrgkelzrbjteisnujiryfqjqkgoqxevtoqlesrlmzpewhxnngdcmdwxaqyinrxsdvtwepoznsvsybbvqdieffigolowmehabgcmpnxjjjaexvhyesisxoeqvsacwmykfvzxudoiedivyfxyeduoghowpyzatvakmfndewyjobxdscamkhohlybdmnsiawzehtrhktmgtgbarisackxsjvekfdlhsgawsnceauavwzmxtljxssonuwiwftvpjeufqyezgyutfvfsqcqtpkaesibkwxkhhvgtwwoebziplwwttuvpmypqdqjiwcamdlwstprplbunwjcbzkjhpgqjvnhcvpnzyrsvtsmcezkfehxdekjwxjeohwuxljqoyrdvziwflwvgdmwiqtumiubqwjfgouddnhuzxgmdjaujxaxnfdjqatgybqznkutkqakzmkcowibruocgzophvbvfsxdraricdltozvxzppxgbckbyiibpnzhwtebjzudwlshwirltmvmjcitwtysccjcjzjabzkvhclwmufgrtbavgknvvefhozasnodnrjernryytmzkmgobcvckeiqahlnwkaippnrqfmzaimoccnssbrmlmoxknwwpmwiblqdntffdjumbiejzexfgvrrbcqwyqybsehytbssbwtovyeencqirqfhaxtylxxnhvksytmbmttaggtijmoywky", 30}, true},
		{"", args{"jicvyghcmoimajglnduqaofybncnhqjltspdmxfayjhkckqzybfrtftlktxriqrhztewahqdwhricmzsbrdfmqlxdqmyizqbbpyrciffjvmyhdsajeltfzvrqhylihmzqafxtzqcwytdwjirprxulapmfsyzxtvyknoyuojsykkpbazofjltgfkqwxevabirswqprbhshytqalxuidbozpunjmjhvvkijqsiqtuwynwkyvmotoprnzdtrltnqpxufjsjqpiyloaaflbuymgscbnlgbieraxxxukaiahgcdzvjrbdhrlvevtxfknexwbflbyblmtkjmhsuimbmnmzleghcsadzluvebninvyecaqemoekuzfdrkkmldymenrgaxgutmxknqtyisumswnqwjedbrfhnpsswazyllymmbpgizoaajgpbpgqtwgzcrzvmppddmzhtmpjkvtxezhbwakikhrevkxevtrhulwfdebxviqsnxsrnrjiklsdnamhetchhkovsqzasqdxqwunddhcrjkxtiavhuhsuzrwjervjihqecoxpsbzbyontpuublomypgkmmobkffeqtcmmwmokqvtaglhihiybsmaueyhfwtowwvetdgipwsnkyvhmwpoehemdysqcsuxlswlhjhoxtgcsbguvyvlcabneiwzxgztcvwnrbofonmmayiyooebjaeuizhxpsolvflvaujlgrdcwsmtpuvfxyxnppnwrwyfqnknavglqxttxemdgyrfizjywhuqikzloydisglmzykkehjzmobpcuzfffgblarallskmzdbcyqhcdjwjmfowunmzosbmbzzxukvfftsfbdejlldetrcwngbrtvesovsyiizqnygvkckdgijvlyanqoyebkefgbgovafwourxinzcytdgfpenglthmmmzmesncotoejiqpmlamzopwndqrxreeksxsidtitcjdllsnrhutulkgqtoxpsgknxxodvhhtgfqeeuxruislmuaueedoqbcvdnqwdshflmayvkcevagvtjynntqgcavooxektxsuxkrhswalvdhwmdffufdsfquoraenkiduonpgjjyqtdqneivfhjfkhddeofstzfshydbcidgvsmqfiqjrmfnxzfdbaqsienfuyejfmcfgqjbedkaomuaqzjvgucudywfahiggvqrzowguyrawliqmuzcymwhkqthyualxpeutpivijyilfqtvuhrynqvdywaexbyfuatywsrxprpywozxiuxsmxpinjvnakrambuofagngtecwfmwuvgyoawzpkyxlrptusjhbrysiisrkqyygvhousljokyrrdelpzzzjmvdajboglfybxehfoyjtobtylkcfuvghdufhsmvgfuhewcyyehdfnbkkuahbbnhyryjrlhuubldduitctbpivimdxosoisxgriaxnyvpokakisxtznyavrssbccvhmgfzlgljpqyoaaqpgecsegtwyvmizfsgvdnhbuqypvogqaagymtewzugdlsikpddqkjozibafmhocrfptfyvwopddlmntcvomsjchhetwqudsekadqtrwvkpljywkshocspreejynixonhdvflpvbqdmllmpysnugglwiiddhxwgivbppybjvusyfrftqunzsqhvlvrddvirfctrrzrkzudhjkclgslzbyfbsqscqvmtnoigkcgxlskdpccagzcccuvliuwmxystoflhbvweqxpliseynangwenzkxlsifktuaehrxhrhrelozibmfmqtbjzbddhnyimnzcsgovljoqxolldcddvcvqvavzyjyhlnvkvhrwdergrgiorjbbjekvtrpqhhcvxeyboslmaygwicyfpnlqqkydwinjzrwttphvnqcjiuwekgqvpwgvnolsbbbkzkaiazdkadvmoibxdwvnyckcqrthydoxqirptwpajtgvnooxozhmulychpgtmamezooplwkaqwyjglkqdnauuximhyltdosuxmfojrljaveozqurksadpfmuewsopdmnijgxqldlxxmhrzymtspvfvvstnvcfofuqvevsakzdpmbyhoayeaggyohzniqbjyvfermsluqzwhzecmganazztsrmakwxcfgswlutpzytxyoriwbhsorocfnqnqnlkvjzmgonocijgwymlolmmddcwcumkjvgssaeuyryzwklgkdxjrvagiwnebhnctaegkvtmiriplkjvqhsysnkjgoucqvfojtjixymhxsntcqpvlvusmvlvojymsmbxvmxderlqnoqxlkacdhduovowkvcglqekndldwgmqyrubmybtpwmoeroyjnglevbtwuejbkihskicqlsfnrkrwgrenawqwwttfpdtfllenuhpywaebqfxoydiamwqbqylcpbbomxodvzpbhwbjqqdnmthxobwnarrdipycxiphvmmkrsclvjaaxwkropsmeyifyephclmlaavurrnlwshlkegouwgbhzmsfwyjalscljjrhgrdijuvxzhvyiqniizilyuaiocxmkrkfkkpfaodauoanfxzxlgxrgfjyxrobskrdwjbbsjnlbcfrmxzzgdkwygxawwtrxwskagltnfgzesobykzjpfcpoydsatvhsckoedwbeztkpoeavlfothdvobuqquefkgnketxlqifcqsrxyhepwdcrncaxnthcxaoopxrwfzppbwfhwgrgvewgescidbwsqjiacsabhpbkuzgqrnxziuxvawztgakeitsahsexgyqqlcmfvsqzmebfwzyycwkmvmvpdesexojbfjvywsuozhlycpshrhxaezjcsfrhhmqfsuffqwdzbqgfyhshnfxmwdhykvmhnxxcjumluszmjylokpxepicmotebyxkgrhkdaunjfpvgmdyorapfvnkrcpysethibkmzbxgatdmksbsoamfwlkdsnqgtmsyajnmouktlulfsdtwjwonaeoaeztrqdmyovjortkpnvrrffugmzxsqfnfavinpiuomdvqbpfmmvwkpxubuftqiphibmkftevrcdornqomrebhclzlgjxszpjegnhuzxfezfmaqcojgfydubevyxcekhrkxxszohvuuyyqshxezhfqebooiioqteawppzgxtfllaepmcqzqmvacpiczulbwozpqolsrwxmtmpmbodyyokiqjupvhdsnyqplozgcidylxbdosfidovutcadtnrugvtyctskbzcpfooqmbbqtvjpekzbvtsznzifdoibhuaeexgtgwqxgpvcbgfmmmjqxgcriguvdinadjbdspiomzakpwuomgtxyouulaywvcdiczonolatslfiaplffwsneeqpealnqswpwcntebbypyvhgbmwhrgalodrifvadulomtmjhfvarlstbdzbzeargwljwgiklmqlopeouediitubhubvoctuumolokkhznmtzmtmmdttsuijcdewuapssbpmqfpltsghzifnctulfmvqbardkuoubdsneydsntuzhjwiomgsagfkcdenvxpbkvmnxhulrwonxdpoocchsylfrvctwytuwdzqlosndffjsleodetmalsnnuhcarhuprywbzzhzddifidwscnpdqlohgkqbxboddzvwrykscisxzzyzumkzjmhtcvzxpzvxmpvudhulkoqqtlnhrlioxhxotgswuqcrkycvigmospqiefibnyileytqkvfbkyaoxnpvnuisogxjzrncqhikfzpvcagjhchmbfcdcqgveagdmfozfxyelnixndlzmoomxkfbcnabnozivyoxmhcughoajpztxyfptkojiybnpxxaxsxwisyimzgvdbyxvfabewuvxhwawdccuhvnjnppsnyeplwuvxdexrviucecpnllpmmpcfcyazvnkxslvdmkzrcierbxcxnahpxuallyujixkwgvyuaiqpvheslzipcjhbhwgvhhotgofrtqfjltbfqnnuggasdmivtbmnhdqestdpdnirydfmizyupymlkpyptdbxsjogpwqknxhgturwuuoqxdnxsxjbfjpbxunogpkkaapmxdxuzmhrewztgagmvxirzwkgikyydkdoztbyslbbsypvnmphzwcjsalvkrsfphbrmzppbegpfftvrkqfwpvqjtkbjlvadpxfnyazufhmgfcsihcmiylaufpwqleizshzgkdtgusrymvrvlpuhneeappvhwfkbgbbbmxzhbeiuiqgdapepnafetgbjrdtcrajqewjiqhtvvthpmkxpxbitgxcqxkysakkdynugeeyhsubhdbaxhgzywhogcqrbprrafdefwujtvlehuhbdaqkxdupcksakcwrppogojbaxnskincjxhgsewnplrwdvmhrxrkvxtmncvqzmpctzdjmgqudfroupeuyqthzmwdwaviydjchryboscstbzjpulnmvktgvgbwncwecjbvyhkfegzlsrexeqedgcezgkdbloulgfonwhykhcdxfnvwnmoqymszmqqoidvabizkbictmpqlhozbjieizpaowspugucteoxvsuwflltfuorydsawjlfoxfcgvojlqurqxbondtubevljbujsfdatlcusmccmhsxagqqprfdcpioygymakztdjwjxovmpdtciyybljqjgydgkcqegeldmlozxlynwtvkpwyqyhiohlbscpbrebbpmxhnnrupxxugqnjqwfeyskefssosjpnbhuwtcqsdvvgqpvsukdfvshxjoioygpxxnpwerqxeskspntkwnrbzmwsxwrvwkdtyhygupbmltsvdpkbphqjhydaceemedkonbbemxkjbmxfgnebhbbxqkzzjegppjikluickomvxtnuxeuparcnitptficmxihrmcvbyhgqshgfjflzcmojfbqiaayuxpqwcpkpfxgwsjpjupwhlmigdqvmronteoukvlvzfhtzfakujlfnszgljoklodvqhxzwmvloqkibblusdrjosfifhcrlfthnrvvmtpqtmuijffdbkcrmjanqglovjggxzltjdxfcxubsyuphhpfvljdgdmfaqoiyqcxggggqrafhaktpljsfssxcszzwzyhdqjpjhfmfatxcsynoeugeievuhhisrylfnxmsbnjusplyqekeodropnnitevippnctkukrkrgdpiitbxezqyxeosrdtbfxdimyrqkywvnpdzciuzqumlozkgqnynkxbiyzfaazdvibjtzhvuooeuxdhnjuudynkuxlqwdnzwakpfzsxxkqnjeragaovyhvqhdsakbuwvvwdqtewjfjnusqfeddxhixtlrpnrjblunqdfzsfhrygrcbjpxfwoymbxymehyivmtgsrtocirskyrdjqwsfsozmyzjsyqeknrrqxtelqpbtcrtxmsqahgxkpxvkoqyncctdfzrzngfxxziplborvxoeqgfoyixppnlrvgdpwiibbfsupqukljrdrsygmqovqhuyhxhwtnyjwddomjxizdxkngpkofxfhqdtwbavqmlnnzjoabzsdfoxkswviwexoqkujhvjizfofpcdctlaohtikrxvjjgtnuqzigshxxqcxkozsbwgyyvbwfauqawasqwiusmpeohiprtptvahfbkhvuvwfvynrpmtrkrhndvwmjcdticabwctzscvfwoorcwzmrmnatxcjnttlcifrjijcqzswtdmpfjfsjsonijfmcujgytsbgsvytfwtfvflsqusmxbfkcrqfbzzneohlixibqvtrujaqgvrdbuagoruwcmqvpsoqpaxglrmijkxyaclcpbfwxeamsrfnjkxobjpkqanewlgifoyaqnisfzzxleqvvdrlavgddcatgpjsmlasrbrxwgndxkfuorkxoiqvvixudosxvasdzrzycvecbhoujqkxryomjvuslodahbetlmlnrphpchcsiqlwqbqhfmobrerigizkczhgxypzbodgjfvwhimjkmjdfskhxnmwfaocsidmfkojerqxedqxfpzcdsoxjgtpjyghruhqayvncuqpvunnxkuntmkejrxlwxomoxmkcimdazavhwlicygroeafblsjehqkqgvaypjhcyumglsroqfdpdvmmppasjnwuludinscdkxigkznnrttezzwuttwwjldvkvphsbxjbbbwinqsiakkalvlclybleneoyxcpfzguoqxmxgqzccwiyveppvdqghfxxncaqsegjqikbxzjscvcwxmjokxbevpwdikukmknkhjeejeikkgsckttmpzcefbvrkrnculppfkmztwdroountdlfqeypkczikjpbzrstrsdwbrbxxyerfvxnslmahyqkxkvhborpprxzdeucqjljwqdzfgpyyjuhywmnclahymmtgdbskqokvecpjlbxyacyvwdrrtitiktgeqncdlgqlmytqprjitfrdchwwamfmoaaigrgkdkxmhbmfiteuqlrwzqnobnonrbqvkwyrshfcxyvwyjrlcmlesljklsqjmvtxrcdruuefhwfskkeadzzozhpznurhexmyrcdvbcndehuzjsigsmerptsjopdbpwfpmaqvagpofzlicvzhhujlxvfbvkipooozbvovaemdvycnnckynwkhmnptmgutivoswhovgkdatsomrfehfatwqtifkjwcumwijovawqbeltravvvxrpzfjznegbwcmcqwzkryarijwlywposiktyfhemtkldxfovhnmzcgavoyusanwrtkqhzgbgqdfdgfitwygsueixmqptsrrwdwmxenlykonnmrfppgdwpyylrlzmqbdhklaolliunjpbfxrwfdfakoltifjvcrnudbgzfvmuqrcgrasvihkvwcujpordbjyeqbossywpwrczlfjvyhmgnbpgnbbtchcbfcqdjomcjmzlsolpcnwswpghrcbbugenyfikalenhntqelsxocdcxpjydfvlskbtzzwzehqvelaqgfoaypkokhhomjygiflvuauqxvrvfhbajczotgqnnhiobfnxhreqxsgnzdqmxbsktuwwzofqazhufkkhztyfjtvpmaqcedsxlkqdemzajnvsczecsffxycufnoxbypdybofclpbaaaohchprlupxvzhsjqkjopnqftievbednmkrcfgqdcotlddbcbxfdlifyeandvisofxkuwckdyimdhpkptufbsshxwskiyslxlajnibuqezzyagizkmlonpkdbsifjswqnkcgovcgdzocsfdpjykxculauqiocksnxdgixgluefbzhdnhzlgstawujlapszfvfemntgbnskubpoyiimliypaibtneayrxwrwwzizgxalfeqohdxjyrrsywqxygvwjgbuseadgorrgikyzsrwkgkoalenpjkokbozzrhzjpcigolnjbvgqwhawjmjrxgilujwteogpiuywijqixiwwckioodrhtljnqcmqjbumxucjmfbhshnfweighcecnbmfrwvdjmlbyltuzbxdcymajolekhuvvpfwaclvhedbtpsvlsrubgegtttabrlzyrdkmkzhkmwlykyxqrqcducvfkmhvsotydejhdqclesgohkbnjyarobhopmunwmhghbgfshcevtgxlqluzxkvgxpsebtqnmqnpjlaqmrhxxqmourisxqcbconlczczxgrlmhpnkcdrynyvnnkgrtvszpqfbfbchxypvzcdkynebypztpoozyeflhvmqcfxhimlcbdxnzatptqhwulysreypoeiyheznvnaavibfedotpkllpebvxjlansxqaqtuwimlvnfaaaqombeamcipnbtaupmhkxonhbdqkwwdjkhokuaoxikjjxwmnwcvuhbzwgnkemwnhmrxlijfyfodldilcictomcivgmdalevudrtdetqhbnixiostbaisqcrklgpzqliuxmamdgujfyvbndnexnjrswsvakwsmisebaolorbhuaiimipygcndhhhdqhvfrkmzezflyrwlhqehwjvzaczakeemayjemkfyttvnoijgdfvhspgvealmxuuvuqrefqvjwweunpeehtzsecxkqauizkbccapqecwxjyngefeuidpcvguintxczhambtnhehnqndbbtxumkcrswkiqpvwgawxhwpfvhywninbamiyqbdmhrbfelmgxpbuytbyzuwbbrhtgrkgcswftipyjjklmuaxysezvwmtcuzwynwgoseheohfrdcvqydzxhcfsfzujoavcbvvdslstxjyendstahtzxozitfnagtfidrflzrvffzqwkohxmsfujcvrffrhqzntrqxdldovzjhunbxlpublccyurynmoqkfqpjnpwkvhemftmdicqppclpvbynvbjlvxhiaqcnupiarkpzgrodohlrgeqberkycramdwxiyivbnlimtkgzpleopkdsswtxipcgacznhdbvkbhxotmojsknmbmxbvhsldetakwnlfjrbjivzzfwphbusvjqmtatzmpseudzolkjlbhmflqbxrsrrbvxnvirpouimaeucfmaruxrokacyxsctpautbrebjwrmyumpjixzhrwumzgvsjjsnbxwptvgstduxsaldgzgveetocbrqmeyfhxxgzvotryqvbebccraehwtbsmeoaywrbkrbskcfwkymnkdjyournljcndjwczzyaqxwzseycgprrxktpfgpofxzrebpxsiyqigltulvlyrtdwqshmurpzwzjyjkordwjmkipwkwqaupwjwqmrwgbbpjvqaxgxhpisshlevjmigpdizsucrjzowhqerzhpoipkzyzxvdjtckvlfyfzflbuqdxkiobwtiacxgvynyumjpprkguikfhpwtvzadjccitvjunyyavgfaetgkdvcjgejkrfpobdkqhcrqucynvwcmeonkmevlbrzabizxionsbahnlmtgptirxtpkrnrphtfabkcbtrfvtjzlaqinvtawpmswjhcnzpmdxzrgvisuiayddrsdkipmnqzifzxcffygfvjbsbvybitcmdlghaomdkaiaivkrtqkslbqfqrfafpqoptlifdfihsaxaxxwqoaqevyhmjbyuqurqgucduwwvajrjuteyrnefpdcozoafnvnnscheiqhelhofrbofqzirfprgrfbrjhuwagqqmvfiacnevojlvyoqtzewhufainogziijfljoojqgqbibixsvgzucpxpnwaxaulpyzxofebjmfygedwbpmwwcpvhmrpcrpnislwnngufzicvzyrivaochnqytdvrwnuzqkwtktzrthoezcoyhfrhkopztajcurkbkacanjugvxxvqryjlymwebfpqmywifdrwfzkxbfcszynhqhjogmtpkpcxdobzjnwibtkrodvgfofkdppzrbdswvmlevkrgvqymlkdiycdlsljlziolhtbpnntkgchjeblknawvpapsgopiraitvnhbyznepfyvytzezzqjkfcnvdbfihetmzfcceokbrumxxeitwwjmcofqxxyhmgstlszqbodfqqkdoaxxdilnvihoksjllqthuzzxugmufhgkydmomiwefbkzivnptrmmhslrdiggbkhpuyinqlxlalqvokdivaueovmryfluvveioyfrsuyiptlzmmlbdtqskeullfcioecskjqrczgbzvqpfkgcuvyqezwafpyxotivuyqjgyeybbxujsrpevkdvdksaspeaidrfejnntjdhzwistzhlhhyezqtthtvszzfyekikwnukvdiwcdsghfgscjfolitmlzregqpcewpfobpqqkrwmzujxlskwmibtqsoyxdlltogwxpzhjoywznhtrmowmapxxfdntozjllhnlryxtjeabhlkqlmmnjxqrozbrajyptfjngytuheoyoaigzqlcpfwbykqkiwtsfkflnccadfvrdcvuhyxyrffhphxzyzxvuhdjejltjqyuzikoyrxuntnhbidtwttpxdtgmjkkdwroihtkxsyhpoteubsxtyvnuioyovvtcbpdyaidjnbsxieduyxgsfrkcfcmgnignlploxhkljgjyfffbmijwclrddokzpazsntvoiqnpykcakcdancglzwtrgohjqwgkiijctvyqxskrafklplqdyqblrhiaavjqk", 4659}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
