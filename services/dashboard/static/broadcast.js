/* Stream-safe read-only view. No prompts, raw logs or private memory are fetched. */
(function (root) {
  'use strict';

  // ---- interface language (same contract as the dashboard: localStorage first,
  // the persisted /api/settings value wins once it arrives).
  const I18N = {
    en: {
      'bc.titleSuffix': 'Live writing room',
      'bc.brand.aria': 'Scriptorium live writing room',
      'bc.tagline': 'Happening now',
      'bc.conn.connecting': 'Connecting to live data',
      'bc.conn.live': 'LIVE · real-time data',
      'bc.conn.paused': 'View paused',
      'bc.conn.reconnecting': 'Waiting to reconnect',
      'bc.select.aria': 'Choose a novel',
      'bc.select.loading': 'Loading books…',
      'bc.select.empty': 'No streamable book yet',
      'bc.pause': 'Pause view',
      'bc.pause.resume': 'Resume view',
      'bc.pause.title': 'Pauses page refresh only; the novel keeps running',
      'bc.clean': 'Clean view',
      'bc.clean.title': 'Hides the controls; press Esc to return',
      'bc.clean.toast': 'Clean view · press Esc to bring the controls back',
      'bc.fullscreen': 'Fullscreen ↗',
      'bc.fullscreen.title': 'Fullscreen · F',
      'bc.fullscreen.embedded': 'The embedded browser did not go fullscreen; use browser fullscreen or an OBS 1920×1080 browser source.',
      'bc.fullscreen.unsupported': 'This browser cannot go fullscreen; use browser fullscreen or OBS window capture.',
      'bc.session.waiting': 'Waiting for connection',
      'bc.hero.title': 'The birth of a novel',
      'bc.hero.desc': 'Connect to a real run and watch character choices become story.',
      'bc.tag.agent': 'Independent character agents',
      'bc.tag.arbiter': 'World-rule arbitration',
      'bc.tag.canon': 'Accepted chapters enter the canon',
      'bc.position.current': 'Current position',
      'bc.position.final': 'Last known position',
      'bc.chapter.unit': 'ch',
      'bc.chapter.waiting': 'Awaiting book data',
      'bc.chapter.total': '{n} chapters total',
      'bc.chapter.totalUnknown': 'Total chapter count unconfirmed',
      'bc.chapter.of': '/ {n} ch',
      'bc.cycle.none': 'No cycle recorded yet',
      'bc.cycle.awaiting': 'Awaiting cycle data',
      'bc.rail.aria': 'Writing stages',
      'bc.section.characters': 'Character activity',
      'bc.section.charactersLive': 'Characters in action',
      'bc.section.charactersLog': 'Character activity log',
      'bc.section.delivery': 'Delivery meter',
      'bc.section.events': 'Live dispatch',
      'bc.section.eventsNote': 'Only real records, never simulated progress',
      'bc.agent.waiting': 'Awaiting character registration',
      'bc.agent.count': '{n} independent characters',
      'bc.hub.idle': 'Awaiting world input',
      'bc.hub.caption': 'Decisions belong to the characters; outcomes answer to the world’s rules.',
      'bc.stage.pending': 'Awaiting stage info',
      'bc.stage.running': 'Executing',
      'bc.stage.awaiting': 'Waiting for the next run',
      'bc.stage.needsAttention': 'Run needs attention',
      'bc.stage.assessing': 'Chapter convergence assessment',
      'bc.stage.preparingPlan': 'Preparing the formal chapter plan',
      'bc.stage.arbitrating': 'World arbitration and validation',
      'bc.stage.deciding': 'Characters deciding independently',
      'bc.stage.unconfirmed': 'Unconfirmed',
      'bc.stage.descRunning': 'Character decisions and world consequences are becoming chapters.',
      'bc.stage.descIdle': 'Not currently running; showing saved real progress.',
      'bc.caption.assess': 'Real consequences decide whether a chapter may enter writing.',
      'bc.caption.check': 'Checking time, resources and conflicts without rewriting intent.',
      'bc.caption.independent': 'Each character acts only on its own observations and memory.',
      'bc.caption.stale': 'Keeping recorded results while waiting for fresh real run data.',
      'bc.empty.characters': 'Registered characters appear here one by one.',
      'bc.empty.charactersLoading': 'Loading character state…',
      'bc.empty.events': 'Waiting for the first run record.',
      'bc.empty.eventsNote': 'Waiting for the first run record; connection heartbeats do not count as progress.',
      'bc.empty.eventsLoading': 'Loading run records…',
      'bc.legend.done': 'Submitted · arbitrated',
      'bc.legend.active': 'In this round',
      'bc.legend.idle': 'Dormant · waiting',
      'bc.privacy': '▣ Stream-safe view · no private character context',
      'bc.metric.accepted': 'Accepted chapters',
      'bc.metric.acceptedCaption': 'Formal acceptance only; simulated cycles do not count as chapters.',
      'bc.metric.planned': 'Chapter plan',
      'bc.metric.words': 'Chapter words',
      'bc.metric.cycles': 'Committed cycles',
      'bc.metric.clock': 'Story clock',
      'bc.usage.label': 'Character simulation usage',
      'bc.usage.recorded': 'Recorded',
      'bc.usage.includesEst': 'Includes estimates',
      'bc.usage.input': 'Input',
      'bc.usage.output': 'Output',
      'bc.usage.calls': 'Calls',
      'bc.usage.unpricedCount': 'unpriced',
      'bc.durable.waiting': 'Awaiting durable progress',
      'bc.durable.noTime': 'No commit time yet',
      'bc.durable.since': 'Last commit {m}m {s}s ago',
      'bc.durable.labelStalled': 'Awaiting a new durable commit',
      'bc.durable.labelRunning': 'Progress counts only durable commits',
      'bc.durable.labelIdle': 'No new work running',
      'bc.footer.motto': 'Every choice leaves a consequence in the world.',
      'bc.footer.shortcuts': 'F fullscreen · Esc exits clean view',
      'bc.refresh.live': 'Read-only connection · updates every 3 s',
      'bc.refresh.paused': 'View paused · the novel keeps running',
      'bc.refresh.lost': 'Connection lost · snapshot kept for {n} s',
      'bc.warn.paused': 'Only the page refresh is paused; the novel’s state is unchanged.',
      'bc.warn.disconnected': 'Connection interrupted; keeping the last snapshot, so state may have changed.',
      'bc.warn.shape': 'Live data shape unavailable; keeping the last snapshot.',
      'bc.warn.connectingBook': 'Connecting to the selected book’s live data…',
      'bc.warn.serviceDown': 'Live service unavailable; reconnecting automatically…',
      'bc.activity.continuing': 'Continuing authorised work',
      'bc.activity.dormant': 'Dormant · waiting for an event',
      'bc.activity.revisionDone': 'Revision decision submitted',
      'bc.activity.revising': 'Revising against conflicting information',
      'bc.activity.submitted': 'Decision submitted this round',
      'bc.activity.independent': 'Observing and deciding independently',
      'bc.activity.needsAttention': 'Awaiting attention',
      'bc.activity.awaitingInput': 'Awaiting this round’s input',
      'bc.activity.lastSubmitted': 'Last decision submitted',
      'bc.activity.resuming': 'Waiting for execution to resume',
      'bc.detail.notRunning': 'No new work running; showing the last saved state.',
      'bc.detail.continuing': 'Reusing the existing authorisation; no extra character-model call.',
      'bc.detail.submitted': 'The decision stands alone, awaiting real results and consequences.',
      'bc.detail.dormant': 'No triggering event, so no extra model call.',
      'bc.detail.private': 'Private observations and memory are never shown on the live view.',
      'bc.character.unnamed': 'Unnamed character',
      'bc.character.calls': '{n} calls',
      'bc.character.input': 'Input {n}',
      'bc.character.canon': 'Canon memory {n} ch',
      'bc.character.round': 'Submitted in R{n}',
      'bc.character.roundUnknown': 'Round unconfirmed',
      'bc.event.chapter': 'Ch {n}',
      'bc.event.cycle': 'Cycle {n}',
      'bc.event.updated': 'Run record updated',
      'bc.event.durable': 'Durable run record',
      'bc.time.unknown': 'Time unconfirmed',
      'bc.book.untitled': 'Untitled novel',
      'bc.book.connecting': 'Connecting to the book…',
      'bc.book.loadingNew': 'Loading the new book',
      'bc.book.awaiting': 'Waiting for a story to begin',
      'bc.book.hint': 'Start a novel and its live status appears here.',
      'bc.book.none': 'No books yet',
      'bc.cost.unpriced': 'Unpriced',
      'bc.cost.approx': '≈ ',
      'bc.cost.estimatedSub': 'Estimated subtotal ',
      'bc.cost.pricedSub': 'Priced subtotal ',
      'bc.error.unavailable': 'Live data temporarily unavailable',
      'bc.error.timeout': 'Live data request timed out',
      'bc.error.schema': 'Unrecognised live data',
    },
    zh: {
      'bc.titleSuffix': '创作直播间',
      'bc.brand.aria': 'Scriptorium 创作直播间',
      'bc.tagline': '故事正在发生',
      'bc.conn.connecting': '连接实时数据',
      'bc.conn.live': 'LIVE · 实时数据',
      'bc.conn.paused': '画面已暂停',
      'bc.conn.reconnecting': '等待重新连接',
      'bc.select.aria': '选择小说',
      'bc.select.loading': '正在读取书目…',
      'bc.select.empty': '暂无可直播的书目',
      'bc.pause': '暂停画面',
      'bc.pause.resume': '恢复画面',
      'bc.pause.title': '只暂停页面刷新，不暂停小说运行',
      'bc.clean': '净屏',
      'bc.clean.title': '隐藏操作区，按 Esc 返回',
      'bc.clean.toast': '净屏模式 · 按 Esc 恢复操作区',
      'bc.fullscreen': '全屏 ↗',
      'bc.fullscreen.title': '全屏 · F',
      'bc.fullscreen.embedded': '内置浏览器未进入全屏；可使用浏览器全屏或 OBS 的 1920×1080 浏览器来源。',
      'bc.fullscreen.unsupported': '当前浏览器不支持页面全屏，可使用浏览器全屏或 OBS 窗口捕获。',
      'bc.session.waiting': '等待连接',
      'bc.hero.title': '一部小说的诞生',
      'bc.hero.desc': '连接真实运行，观察角色选择如何成为故事。',
      'bc.tag.agent': '独立角色 Agent',
      'bc.tag.arbiter': '世界规则裁决',
      'bc.tag.canon': '正文验收后入正史',
      'bc.position.current': '当前推进',
      'bc.position.final': '最后工作位置',
      'bc.chapter.unit': '章',
      'bc.chapter.waiting': '等待书目数据',
      'bc.chapter.total': '全书 {n} 章',
      'bc.chapter.totalUnknown': '全书章数待确认',
      'bc.chapter.of': '/ {n} 章',
      'bc.cycle.none': '尚无周期记录',
      'bc.cycle.awaiting': '等待周期数据',
      'bc.rail.aria': '创作阶段',
      'bc.section.characters': '角色活动',
      'bc.section.charactersLive': '角色正在行动',
      'bc.section.charactersLog': '角色活动记录',
      'bc.section.delivery': '交付仪表',
      'bc.section.events': '现场播报',
      'bc.section.eventsNote': '只播报实际记录，不模拟进度',
      'bc.agent.waiting': '等待角色注册',
      'bc.agent.count': '{n} 名独立角色',
      'bc.hub.idle': '等待世界输入',
      'bc.hub.caption': '决定属于角色，结果接受世界规则检验。',
      'bc.stage.pending': '等待阶段信息',
      'bc.stage.running': '执行中',
      'bc.stage.awaiting': '等待下一次执行',
      'bc.stage.needsAttention': '运行需要处理',
      'bc.stage.assessing': '章节收束评估',
      'bc.stage.preparingPlan': '准备正式章节计划',
      'bc.stage.arbitrating': '世界裁决与验证',
      'bc.stage.deciding': '角色独立决策中',
      'bc.stage.unconfirmed': '待确认',
      'bc.stage.descRunning': '角色决策与世界后果，正在成为章节。',
      'bc.stage.descIdle': '当前未运行，显示已保存的真实进展。',
      'bc.caption.assess': '以真实后果判断章节能否进入写作。',
      'bc.caption.check': '检查时间、资源和冲突，不改写角色的意图。',
      'bc.caption.independent': '每个角色只依据自己的观察与记忆作出选择。',
      'bc.caption.stale': '保留已记录的结果，等待新的真实运行数据。',
      'bc.empty.characters': '角色注册后，将在这里逐一出现。',
      'bc.empty.charactersLoading': '正在读取角色状态…',
      'bc.empty.events': '等待第一条运行记录。',
      'bc.empty.eventsNote': '等待第一条运行记录；连接心跳不算业务进展。',
      'bc.empty.eventsLoading': '正在读取运行记录…',
      'bc.legend.done': '已提交 / 已裁决',
      'bc.legend.active': '本轮处理',
      'bc.legend.idle': '休眠 / 等待',
      'bc.privacy': '▣ 直播安全视图 · 不展示角色私有上下文',
      'bc.metric.accepted': '已验收正文',
      'bc.metric.acceptedCaption': '以正式验收为准，不把模拟周期算成章节。',
      'bc.metric.planned': '正式章节计划',
      'bc.metric.words': '正文字数',
      'bc.metric.cycles': '已落盘周期',
      'bc.metric.clock': '故事时钟',
      'bc.usage.label': '角色推演用量',
      'bc.usage.recorded': '已记录',
      'bc.usage.includesEst': '含估算',
      'bc.usage.input': '输入',
      'bc.usage.output': '输出',
      'bc.usage.calls': '调用',
      'bc.usage.unpricedCount': '次未计价',
      'bc.durable.waiting': '等待持久进展',
      'bc.durable.noTime': '尚无提交时间',
      'bc.durable.since': '上次提交 {m}分{s}秒前',
      'bc.durable.labelStalled': '等待新的持久提交',
      'bc.durable.labelRunning': '只以持久提交计进展',
      'bc.durable.labelIdle': '当前未执行新任务',
      'bc.footer.motto': '每个选择，都要在世界里留下后果。',
      'bc.footer.shortcuts': 'F 全屏 · Esc 退出净屏',
      'bc.refresh.live': '只读连接 · 每 3 秒更新',
      'bc.refresh.paused': '画面已暂停 · 不影响小说运行',
      'bc.refresh.lost': '连接中断 · 快照已保留 {n} 秒',
      'bc.warn.paused': '仅暂停页面刷新，不更改小说执行状态。',
      'bc.warn.disconnected': '连接暂时中断，保留最后一次快照；状态可能已变化。',
      'bc.warn.shape': '直播数据格式暂不可用，保留最后一次快照。',
      'bc.warn.connectingBook': '正在连接所选书目的实时数据…',
      'bc.warn.serviceDown': '直播服务暂不可用，正在自动重连…',
      'bc.activity.continuing': '继续已授权的工作',
      'bc.activity.dormant': '休眠 · 等待新事件',
      'bc.activity.revisionDone': '修订决定已提交',
      'bc.activity.revising': '根据冲突信息修订',
      'bc.activity.submitted': '本轮决定已提交',
      'bc.activity.independent': '独立观察与决策',
      'bc.activity.needsAttention': '等待处理',
      'bc.activity.awaitingInput': '等待本轮输入',
      'bc.activity.lastSubmitted': '上次决定已提交',
      'bc.activity.resuming': '等待恢复执行',
      'bc.detail.notRunning': '当前未执行新任务，显示最后保存的状态。',
      'bc.detail.continuing': '沿用原有授权，不重复调用角色模型。',
      'bc.detail.submitted': '决定保持独立，等待真实结果与后果。',
      'bc.detail.dormant': '没有触发事件，不额外消耗模型调用。',
      'bc.detail.private': '私有观察与记忆不会在直播页面展示。',
      'bc.character.unnamed': '未命名角色',
      'bc.character.calls': '{n} 次调用',
      'bc.character.input': '输入 {n}',
      'bc.character.canon': '正式记忆 {n} 章',
      'bc.character.round': '提交轮次 R{n}',
      'bc.character.roundUnknown': '轮次待确认',
      'bc.event.chapter': '第 {n} 章',
      'bc.event.cycle': '周期 {n}',
      'bc.event.updated': '运行记录已更新',
      'bc.event.durable': '持久运行记录',
      'bc.time.unknown': '时间待确认',
      'bc.book.untitled': '未命名小说',
      'bc.book.connecting': '正在连接书目…',
      'bc.book.loadingNew': '正在读取新书目',
      'bc.book.awaiting': '等待故事启程',
      'bc.book.hint': '启动一本小说后，运行状态会出现在这里。',
      'bc.book.none': '暂无书目',
      'bc.cost.unpriced': '未计价',
      'bc.cost.approx': '约 ',
      'bc.cost.estimatedSub': '估算小计 ',
      'bc.cost.pricedSub': '已计价小计 ',
      'bc.error.unavailable': '直播数据暂不可用',
      'bc.error.timeout': '直播数据请求超时',
      'bc.error.schema': '无法识别直播数据',
    },
  };
  let LANG = 'en';
  const locale = () => (LANG === 'zh' ? 'zh-CN' : 'en-GB');
  const t = (key, vars) => {
    const table = I18N[LANG] || I18N.en;
    let value = Object.hasOwn(table, key) ? table[key] : (I18N.en[key] ?? key);
    if (vars) for (const name in vars) value = value.split('{' + name + '}').join(String(vars[name]));
    return value;
  };
  const setLang = value => { LANG = value === 'zh' ? 'zh' : 'en'; return LANG; };
  const getLang = () => LANG;

  const STAGES = {
    en: {
      architect: 'Building the story world', 'outline-all': 'Freezing the book outline',
      'zero-init': 'Initializing the world', preplan: 'Preparing the current story arc',
      'rehearse-arc': 'Rehearsing arc contingencies', 'project-all': 'Projecting each character',
      seal: 'Sealing the chapter plan', promote: 'Preparing the chapter contract',
      render: 'Writing and reviewing chapters', finalize: 'Final review of the whole book',
      deliver: 'Assembling the final delivery',
    },
    zh: {
      architect: '构建故事世界', 'outline-all': '冻结全书导航', 'zero-init': '初始化世界',
      preplan: '准备当前故事弧', 'rehearse-arc': '整弧条件预演', 'project-all': '角色独立细推',
      seal: '封存章节计划', promote: '准备正文合同', render: '正文创作与审核',
      finalize: '全书终审', deliver: '整理最终交付',
    },
  };
  const STATES = {
    en: {
      running: 'Running', idle: 'Standing by', paused: 'Paused', error: 'Needs attention',
      stopped: 'Not running', complete: 'This run finished', completed: 'This run finished',
      attention: 'Needs attention', unknown: 'Status unconfirmed', waiting: 'Waiting',
    },
    zh: {
      running: '运行中', idle: '待命', paused: '已暂停', error: '需要处理',
      stopped: '未运行', complete: '本次执行完成', completed: '本次执行完成',
      attention: '需要关注', unknown: '状态待确认', waiting: '等待中',
    },
  };
  const OUTCOMES = {
    en: {
      success: 'Action succeeded', partial: 'Partially done', blocked: 'Action blocked',
      failure: 'Action failed', failed: 'Action failed', infeasible: 'Hard constraint conflict',
    },
    zh: {
      success: '行动成功', partial: '部分完成', blocked: '行动受阻',
      failure: '行动失败', failed: '行动失败', infeasible: '存在硬约束冲突',
    },
  };
  const EVENTS = {
    en: {
      proposal_submitted: 'Character decision submitted', decision_submitted: 'Character decision submitted',
      arbitration_submitted: 'World arbitration submitted', arbitration_finalized: 'World arbitration committed',
      cycle_committed: 'World cycle committed', readiness_committed: 'Chapter readiness submitted',
      bundle_committed: 'Chapter plan saved', chapter_accepted: 'Chapter formally accepted',
      chapter_completed: 'Chapter completion record updated', context_bound: 'Round input bound',
      generation_started: 'Current plan started', stage_started: 'Entered a new execution stage',
      stage_completed: 'Execution stage completed', activation: 'Character received an activation event',
      proposal: 'Character decision submitted', arbitration: 'World arbitration submitted',
      cycle: 'World cycle committed', readiness: 'Chapter readiness submitted',
      plan: 'Chapter plan saved', accepted: 'Chapter formally accepted',
      proposal_committed: 'Character decision submitted', arbitration_committed: 'World arbitration submitted',
      stage_failed: 'Execution stage needs attention', error: 'Run needs attention',
      failed: 'Run needs attention',
    },
    zh: {
      proposal_submitted: '角色决定已提交', decision_submitted: '角色决定已提交',
      arbitration_submitted: '世界裁决已提交', arbitration_finalized: '世界裁决已落盘',
      cycle_committed: '世界周期已落盘', readiness_committed: '章节评估已提交',
      bundle_committed: '正式章节计划已保存', chapter_accepted: '正文已正式验收',
      chapter_completed: '章节完成记录已更新', context_bound: '本轮输入已绑定',
      generation_started: '开始当前规划', stage_started: '进入新执行阶段',
      stage_completed: '执行阶段已完成', activation: '角色收到激活事件',
      proposal: '角色决定已提交', arbitration: '世界裁决已提交', cycle: '世界周期已落盘',
      readiness: '章节评估已提交', plan: '正式章节计划已保存', accepted: '正文已正式验收',
      proposal_committed: '角色决定已提交', arbitration_committed: '世界裁决已提交',
      stage_failed: '执行阶段需要处理', error: '运行需要处理', failed: '运行需要处理',
    },
  };
  const TIERS = {
    en: {
      protagonist: 'Protagonist · PERSISTENT AGENT', core: 'Core character · AGENT',
      supporting: 'Supporting character · AGENT', important: 'Important character · AGENT',
      default: 'Independent character · AGENT',
    },
    zh: {
      protagonist: '主角 · PERSISTENT AGENT', core: '核心角色 · AGENT',
      supporting: '重要配角 · AGENT', important: '重要角色 · AGENT',
      default: '独立角色 · AGENT',
    },
  };
  const RAIL = {
    en: [['world', 'World setup'], ['outline', 'Book outline'], ['rehearsal', 'Arc rehearsal'],
      ['decisions', 'Character decisions'], ['arbiter', 'World arbitration'],
      ['planning', 'Chapter sealing'], ['writing', 'Chapter acceptance']],
    zh: [['world', '世界初始化'], ['outline', '全书导航'], ['rehearsal', '整弧预演'],
      ['decisions', '角色决策'], ['arbiter', '世界裁决'], ['planning', '章节封存'],
      ['writing', '正文验收']],
  };

  const finite = value => typeof value === 'number' && Number.isFinite(value) && value >= 0 ? value : null;
  const dict = value => value && typeof value === 'object' && !Array.isArray(value) ? value : {};
  const escapeHTML = value => String(value ?? '').replace(/[&<>"']/g, char => ({'&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;'}[char]));
  function formatNumber(value, compact = false) {
    value = finite(value);
    if (value === null) return '—';
    if (compact && value >= 1000000) return (value / 1000000).toFixed(2).replace(/0+$/, '').replace(/\.$/, '') + 'M';
    if (compact && value >= 10000) return (value / 1000).toFixed(1).replace(/\.0$/, '') + 'K';
    return new Intl.NumberFormat(locale(), {maximumFractionDigits: 1}).format(value);
  }
  function formatCost(usage) {
    usage = dict(usage);
    const amount = finite(usage.cost_usd), source = usage.cost_source;
    if (amount === null || !['reported', 'estimated', 'partial'].includes(source)) return t('bc.cost.unpriced');
    const money = amount > 0 && amount < 0.0001 ? '<$0.0001' : '$' + amount.toFixed(amount < 0.01 && amount > 0 ? 4 : 2);
    return (source === 'estimated' ? t('bc.cost.approx') : source === 'partial'
      ? (finite(usage.estimated_calls) > 0 ? t('bc.cost.estimatedSub') : t('bc.cost.pricedSub')) : '') + money;
  }
  function stageLabel(stage) { return STAGES[LANG][stage] || t('bc.stage.pending'); }
  function activityLabel(character) {
    character = dict(character);
    if (character.state === 'continuing') return t('bc.activity.continuing');
    if (character.state === 'sleeping' || character.state === 'dormant') return t('bc.activity.dormant');
    if (character.state === 'revising' || character.state === 'revision') return character.submitted === true ? t('bc.activity.revisionDone') : t('bc.activity.revising');
    if (character.submitted === true || character.state === 'submitted') return t('bc.activity.submitted');
    if (character.state === 'proposing' || character.state === 'active') return t('bc.activity.independent');
    if (character.state === 'error' || character.state === 'blocked') return t('bc.activity.needsAttention');
    return t('bc.activity.awaitingInput');
  }
  function buildView(snapshot) {
    snapshot = dict(snapshot);
    const chapter = dict(snapshot.chapter), planning = dict(snapshot.planning), execution = dict(snapshot.execution);
    const characters = (Array.isArray(snapshot.characters) ? snapshot.characters : []).filter(x => x && typeof x === 'object');
    const isRunning = snapshot.status === 'running';
    const allSubmitted = characters.length > 0 && characters.every(c => c.submitted === true || ['sleeping', 'dormant', 'continuing', 'submitted'].includes(c.state));
    let rail = 'world', hub = stageLabel(snapshot.stage), caption = t('bc.hub.caption');
    if (snapshot.stage === 'outline-all') rail = 'outline';
    if (['preplan', 'rehearse-arc'].includes(snapshot.stage)) rail = 'rehearsal';
    if (snapshot.stage === 'project-all') {
      rail = execution.phase === 'assessing' || execution.phase === 'ready' ? 'planning' : allSubmitted ? 'arbiter' : 'decisions';
      hub = execution.phase === 'assessing' ? t('bc.stage.assessing') : execution.phase === 'ready' ? t('bc.stage.preparingPlan') : allSubmitted ? t('bc.stage.arbitrating') : t('bc.stage.deciding');
      caption = execution.phase === 'assessing' ? t('bc.caption.assess') : allSubmitted ? t('bc.caption.check') : t('bc.caption.independent');
    }
    if (['seal', 'promote'].includes(snapshot.stage)) rail = 'planning';
    if (['render', 'finalize', 'deliver'].includes(snapshot.stage)) rail = 'writing';
    if (!isRunning) { hub = snapshot.status === 'error' ? t('bc.stage.needsAttention') : t('bc.stage.awaiting'); caption = t('bc.caption.stale'); }
    const completed = finite(chapter.completed), total = finite(chapter.total);
    return {snapshot, chapter, planning, execution, characters, isRunning, rail, hub, caption,
      acceptedPercent: completed !== null && total !== null && total > 0 ? Math.min(100, completed / total * 100) : 0};
  }
  function eventTime(value) {
    if (!value) return t('bc.time.unknown');
    const d = new Date(typeof value === 'number' ? value * 1000 : value);
    return Number.isFinite(d.getTime()) ? new Intl.DateTimeFormat(locale(), {hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false}).format(d) : t('bc.time.unknown');
  }
  function renderSnapshot(snapshot) {
    const v = buildView(snapshot), e = escapeHTML, tiers = TIERS[LANG], outcomes = OUTCOMES[LANG], events = EVENTS[LANG];
    const cards = v.characters.map((c, index) => {
      const color = ['', 'amber', 'blue'][index % 3], name = typeof c.name === 'string' && c.name ? c.name : t('bc.character.unnamed');
      const active = v.isRunning && !['sleeping', 'dormant', 'idle', 'unknown'].includes(c.state);
      const chapterNo = finite(c.memory_chapter), round = finite(c.round);
      const actorUsage = dict(c.usage);
      const activity = !v.isRunning && ['active', 'proposing', 'revising', 'revision'].includes(c.state) ? (c.submitted ? t('bc.activity.lastSubmitted') : t('bc.activity.resuming')) : activityLabel(c);
      const detail = !v.isRunning ? t('bc.detail.notRunning') : c.state === 'continuing' ? t('bc.detail.continuing') : c.submitted === true ? t('bc.detail.submitted') : c.state === 'sleeping' ? t('bc.detail.dormant') : t('bc.detail.private');
      return `<article class="character-card ${color}${active ? ' is-active' : ''}" aria-label="${e(name)} ${e(activity)}"><div class="character-top"><div class="avatar">${e(Array.from(name)[0])}</div><div><div class="character-name">${e(name)}</div><div class="character-tier">${e(tiers[c.tier] || tiers.default)}</div></div><i class="status-indicator"></i></div><div class="character-activity">${e(activity)}</div><p class="character-detail">${e(detail)}</p><div class="character-usage"><span>${e(t('bc.character.calls', {n: formatNumber(actorUsage.calls)}))}</span><span>${e(t('bc.character.input', {n: formatNumber(actorUsage.input_tokens, true)}))}</span></div>${c.outcome && outcomes[c.outcome] ? `<span class="character-tag">${e(outcomes[c.outcome])}</span>` : ''}<div class="character-meta"><span>${e(t('bc.character.canon', {n: chapterNo === null ? '—' : chapterNo}))}</span><span>${round === null ? t('bc.character.roundUnknown') : e(t('bc.character.round', {n: round}))}</span></div></article>`;
    }).join('') || `<div class="empty-state">${t('bc.empty.characters')}</div>`;
    const eventList = (Array.isArray(snapshot?.events) ? snapshot.events : []).filter(x => x && typeof x === 'object').slice(0, 16).map(event => {
      const meta = [event.stage && STAGES[LANG][event.stage] ? STAGES[LANG][event.stage] : '', finite(event.chapter) ? t('bc.event.chapter', {n: event.chapter}) : '', finite(event.cycle) ? t('bc.event.cycle', {n: event.cycle}) : '', finite(event.round) ? `R${event.round}` : ''].filter(Boolean).join(' · ');
      return `<article class="event-card" data-event-id="${e(event.id)}"><div class="event-time">${e(eventTime(event.time))}</div><div class="event-title"><i></i>${e(events[event.kind] || t('bc.event.updated'))}</div><div class="event-meta">${e(meta || t('bc.event.durable'))}</div></article>`;
    }).join('') || `<div class="empty-state">${t('bc.empty.eventsNote')}</div>`;
    const rail = RAIL[LANG].map(([key, label], i) => `<div class="stage-step${v.rail === key && v.isRunning ? ' current' : v.rail === key && STAGES[LANG][v.snapshot.stage] ? ' halted' : ''}"${v.rail === key && v.isRunning ? ' aria-current="step"' : ''}><span>${String(i + 1).padStart(2, '0')}</span>${label}</div>`).join('');
    return {cards, events: eventList, rail};
  }

  // One request at a time. Stopping or switching books invalidates late replies.
  function createPoller({load, onData, onError = () => {}, interval = 3000,
    setTimer = setTimeout, clearTimer = clearTimeout, controllerFactory = () => new AbortController()}) {
    let running = false, generation = 0, active = null, timer = null;
    async function tick() {
      if (!running || active) return;
      const request = {controller: controllerFactory(), generation}; active = request;
      try {
        const value = await load(request.controller.signal);
        if (running && request.generation === generation) onData(value);
      } catch (error) {
        if (running && request.generation === generation && error?.name !== 'AbortError') onError(error);
      } finally {
        if (active === request) {
          active = null;
          if (running) timer = setTimer(tick, interval);
        }
      }
    }
    return {
      start() { if (!running) { running = true; tick(); } },
      stop() { running = false; generation++; if (timer !== null) clearTimer(timer); timer = null; const request = active; active = null; request?.controller.abort(); },
      refresh() { if (timer !== null) clearTimer(timer); timer = null; tick(); },
      isRunning() { return running; },
    };
  }
  root.BroadcastView = {escapeHTML, formatNumber, formatCost, stageLabel, activityLabel, buildView, renderSnapshot, createPoller, t, setLang, getLang, I18N};
  if (typeof module !== 'undefined' && module.exports) module.exports = root.BroadcastView;
  if (typeof document === 'undefined') return;

  const $ = id => document.getElementById(id);
  let selected = '', paused = false, connected = false, lastSnapshot = null, lastSuccess = 0, toastTimer;
  let connOwned = false, warningKey = 'bc.warn.disconnected';
  const text = (id, value) => { $(id).textContent = value; };
  const html = (id, value) => { if ($(id).innerHTML !== value) $(id).innerHTML = value; };
  async function getJSON(url, signal) {
    const controller = new AbortController(); let timedOut = false;
    const cancel = () => controller.abort();
    if (signal?.aborted) cancel(); else signal?.addEventListener('abort', cancel, {once: true});
    const timeout = setTimeout(() => { timedOut = true; cancel(); }, 8000);
    try {
      const response = await fetch(url, {signal: controller.signal, cache: 'no-store', headers: {'Accept': 'application/json'}});
      if (!response.ok) { const error = new Error(t('bc.error.unavailable')); error.status = response.status; throw error; }
      return await response.json();
    } catch (error) {
      if (timedOut) throw new Error(t('bc.error.timeout'));
      throw error;
    } finally { clearTimeout(timeout); signal?.removeEventListener('abort', cancel); }
  }
  function renderConnectionLabel() {
    const state = paused ? 'paused' : connected ? 'live' : 'reconnecting';
    $('connection').classList.toggle('offline', !connected || paused);
    document.body.classList.toggle('is-disconnected', !connected || paused);
    $('connection').querySelector('span').textContent = t('bc.conn.' + state);
    $('connection-warning').hidden = connected && !paused;
    if (!$('connection-warning').hidden) text('connection-warning', paused ? t('bc.warn.paused') : t(warningKey));
  }
  function setConnection(ok, messageKey) {
    connected = ok; connOwned = true;
    warningKey = messageKey || 'bc.warn.disconnected';
    renderConnectionLabel();
  }
  function draw(snapshot) {
    if (!snapshot || snapshot.schema !== 'broadcast.v1') throw new Error(t('bc.error.schema'));
    lastSnapshot = snapshot; lastSuccess = Date.now();
    const v = buildView(snapshot), markup = renderSnapshot(snapshot), usage = dict(snapshot.usage);
    const states = STATES[LANG];
    text('book-title', snapshot.title || t('bc.book.untitled'));
    document.title = `${snapshot.title || 'Scriptorium'} · ${t('bc.titleSuffix')}`;
    text('session-status', states[snapshot.status] || states.unknown);
    text('stage-description', `${stageLabel(snapshot.stage)} · ${v.isRunning ? t('bc.stage.descRunning') : t('bc.stage.descIdle')}`);
    text('position-label', v.isRunning ? t('bc.position.current') : t('bc.position.final'));
    text('characters-heading', v.isRunning ? t('bc.section.charactersLive') : t('bc.section.charactersLog'));
    text('chapter-number', formatNumber(v.chapter.current));
    text('chapter-total', finite(v.chapter.total) ? t('bc.chapter.total', {n: formatNumber(v.chapter.total)}) : t('bc.chapter.totalUnknown'));
    text('cycle-position', finite(v.planning.cycle) ? `CYCLE ${String(v.planning.cycle).padStart(2, '0')}${finite(v.planning.round) ? ` / R${v.planning.round}` : ''}` : t('bc.cycle.none'));
    text('hub-title', v.hub); text('hub-caption', v.caption);
    text('hub-state', v.isRunning ? t('bc.stage.running') : states[snapshot.status] || t('bc.stage.unconfirmed'));
    $('world-hub').classList.toggle('busy', v.isRunning);
    $('world-hub').classList.toggle('needs-attention', ['error', 'attention'].includes(snapshot.status));
    html('stage-rail', markup.rail); html('character-grid', markup.cards); html('event-list', markup.events);
    text('agent-count', v.characters.length ? t('bc.agent.count', {n: v.characters.length}) : t('bc.agent.waiting'));
    text('accepted-number', formatNumber(v.chapter.completed));
    text('accepted-total', finite(v.chapter.total) ? t('bc.chapter.of', {n: formatNumber(v.chapter.total)}) : t('bc.chapter.unit'));
    $('accepted-bar').style.width = `${v.acceptedPercent}%`;
    text('planned-number', formatNumber(v.planning.planned_chapters));
    text('words-number', formatNumber(v.chapter.words, true));
    text('committed-number', formatNumber(v.planning.completed_cycles));
    text('story-clock', finite(v.planning.story_minutes) !== null ? `T+${formatNumber(v.planning.story_minutes)}′` : '—');
    const input = finite(usage.input_tokens), output = finite(usage.output_tokens);
    text('usage-tokens', input !== null && output !== null ? formatNumber(input + output, true) : '—');
    text('input-tokens', t('bc.usage.input') + ' ' + formatNumber(input, true));
    text('output-tokens', t('bc.usage.output') + ' ' + formatNumber(output, true));
    text('usage-calls', `${t('bc.usage.calls')} ${formatNumber(usage.calls)}${finite(usage.unknown_calls) > 0 ? ` · ${formatNumber(usage.unknown_calls)} ${t('bc.usage.unpricedCount')}` : ''}`);
    text('usage-cost', formatCost(usage));
    text('usage-scope', usage.cost_source === 'estimated' || finite(usage.estimated_calls) > 0 ? t('bc.usage.includesEst') : t('bc.usage.recorded'));
    $('durable-dot').style.background = v.execution.stalled ? 'var(--amber)' : 'var(--mint)';
    text('durable-label', v.execution.stalled ? t('bc.durable.labelStalled') : v.isRunning ? t('bc.durable.labelRunning') : t('bc.durable.labelIdle'));
    setConnection(true); tickClock();
  }
  const poller = createPoller({
    load: signal => getJSON(`/api/novels/${encodeURIComponent(selected)}/broadcast`, signal),
    onData: snapshot => { try { draw(snapshot); } catch (_) { setConnection(false, 'bc.warn.shape'); } },
    onError: error => { setConnection(false); if (error?.status === 404) { poller.stop(); discover(); } },
  });
  function showToast(message) { clearTimeout(toastTimer); text('toast', message); $('toast').classList.add('show'); toastTimer = setTimeout(() => $('toast').classList.remove('show'), 2200); }
  function renderBooks() {
    if (!bookList) return;
    $('book-select').replaceChildren();
    if (!bookList.length) {
      const option = document.createElement('option'); option.textContent = t('bc.select.empty'); $('book-select').append(option);
      return;
    }
    bookList.forEach(n => {
      const option = document.createElement('option'); option.value = n.id;
      option.textContent = n.title || t('bc.book.untitled'); $('book-select').append(option);
    });
    if (selected) $('book-select').value = selected;
  }
  function switchBook(id) {
    poller.stop(); selected = id; lastSnapshot = null; lastSuccess = 0;
    const url = new URL(location.href); url.searchParams.set('novel', id); history.replaceState(null, '', url);
    text('book-title', t('bc.book.connecting')); html('character-grid', `<div class="empty-state">${t('bc.empty.charactersLoading')}</div>`);
    html('event-list', `<div class="empty-state">${t('bc.empty.eventsLoading')}</div>`);
    ['chapter-number', 'accepted-number', 'planned-number', 'words-number', 'committed-number', 'story-clock', 'usage-tokens'].forEach(id => text(id, '—'));
    text('hub-title', t('bc.book.loadingNew')); text('cycle-position', t('bc.cycle.awaiting')); text('durable-time', t('bc.durable.noTime'));
    $('accepted-bar').style.width = '0%'; html('stage-rail', '');
    setConnection(false, 'bc.warn.connectingBook');
    if (!paused && !document.hidden) poller.start();
  }
  async function discover() {
    try {
      const data = await getJSON('/api/broadcast');
      const novels = (Array.isArray(data.novels) ? data.novels : []).filter(n => n && typeof n.id === 'string');
      bookList = novels;
      if (!novels.length) {
        renderBooks();
        text('book-title', t('bc.book.awaiting')); text('stage-description', t('bc.book.hint'));
        setConnection(true); text('session-status', t('bc.book.none')); setTimeout(discover, 5000); return;
      }
      renderBooks();
      const requested = new URL(location.href).searchParams.get('novel');
      const chosen = novels.find(n => n.id === requested) || novels.find(n => n.status === 'running') || novels[0];
      $('book-select').value = chosen.id; switchBook(chosen.id);
    } catch (_) { setConnection(false, 'bc.warn.serviceDown'); setTimeout(discover, 5000); }
  }
  function tickClock() {
    text('wall-clock', new Intl.DateTimeFormat(locale(), {hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false}).format(new Date()));
    if (!lastSnapshot) return;
    const value = lastSnapshot.execution?.last_progress_at;
    const parsed = value ? new Date(typeof value === 'number' ? value * 1000 : value).getTime() : NaN;
    const seconds = Number.isFinite(parsed) ? Math.max(0, Math.floor((Date.now() - parsed) / 1000)) : null;
    text('durable-time', seconds === null ? t('bc.durable.noTime') : t('bc.durable.since', {m: Math.floor(seconds / 60), s: String(seconds % 60).padStart(2, '0')}));
    text('refresh-state', paused ? t('bc.refresh.paused') : connected ? t('bc.refresh.live') : t('bc.refresh.lost', {n: Math.max(0, Math.floor((Date.now() - lastSuccess) / 1000))}));
  }
  function applyLang(next) {
    setLang(next);
    document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
    document.querySelectorAll('[data-i18n]').forEach(el => { el.textContent = t(el.dataset.i18n); });
    document.querySelectorAll('[data-i18n-title]').forEach(el => { el.title = t(el.dataset.i18nTitle); });
    document.querySelectorAll('[data-i18n-aria]').forEach(el => { el.setAttribute('aria-label', t(el.dataset.i18nAria)); });
    document.title = `${(lastSnapshot && lastSnapshot.title) || 'Scriptorium'} · ${t('bc.titleSuffix')}`;
    renderBooks();
    if (!lastSnapshot && bookList && !bookList.length) {
      text('book-title', t('bc.book.awaiting'));
      text('stage-description', t('bc.book.hint'));
      text('session-status', t('bc.book.none'));
    } else if (!lastSnapshot && bookList && bookList.length && selected) {
      text('book-title', t('bc.book.connecting'));
    }
    // Re-render whatever is on screen without altering live connection state.
    const wasConnected = connected;
    if (lastSnapshot) { try { draw(lastSnapshot); } catch (_) { /* keep the last good view */ } connected = wasConnected; }
    if (connOwned) renderConnectionLabel();
    if (paused) text('pause-button', t('bc.pause.resume'));
    tickClock();
  }
  function initLanguage() {
    let saved = 'en';
    try { saved = localStorage.getItem('ns-lang') || 'en'; } catch { /* storage blocked */ }
    applyLang(saved === 'zh' ? 'zh' : 'en');
    fetch('/api/settings', {cache: 'no-store'})
      .then(r => (r.ok ? r.json() : null))
      .then(s => { if (s && s.language && s.language !== LANG) applyLang(s.language); })
      .catch(() => { /* keep the current language if settings are unreachable */ });
  }
  async function toggleFullscreen() {
    try {
      if (document.fullscreenElement) await document.exitFullscreen();
      else {
        if (!document.fullscreenEnabled || typeof document.documentElement.requestFullscreen !== 'function') throw new Error('unsupported');
        await document.documentElement.requestFullscreen();
        if (!document.fullscreenElement) showToast(t('bc.fullscreen.embedded'));
      }
    }
    catch (_) { showToast(t('bc.fullscreen.unsupported')); }
  }
  let bookList = null;
  $('book-select').addEventListener('change', event => switchBook(event.target.value));
  $('pause-button').addEventListener('click', () => { paused = !paused; text('pause-button', paused ? t('bc.pause.resume') : t('bc.pause')); if (paused) poller.stop(); else if (selected && !document.hidden) poller.start(); if (connOwned) renderConnectionLabel(); tickClock(); });
  $('clean-button').addEventListener('click', () => { document.body.classList.add('clean'); showToast(t('bc.clean.toast')); });
  $('fullscreen-button').addEventListener('click', toggleFullscreen);
  document.addEventListener('keydown', event => { if (/^(INPUT|SELECT|TEXTAREA)$/.test(event.target.tagName)) return; if (event.key === 'Escape') document.body.classList.remove('clean'); if (event.key.toLowerCase() === 'f') toggleFullscreen(); });
  document.addEventListener('visibilitychange', () => { if (document.hidden) poller.stop(); else if (!paused && selected) poller.start(); });
  window.addEventListener('pagehide', () => poller.stop());
  initLanguage();
  tickClock(); setInterval(tickClock, 1000); discover();
})(typeof globalThis !== 'undefined' ? globalThis : this);
