// DO NOT EDIT. This file is automatically generated.

export interface ChangeList {
	system: string;
	id: string;
	owner: string;
	status: string;
	subject: string;
	updated: string;
	url: string;
}

export interface TryJob {
	id: string;
	name: string;
	updated: string;
	system: string;
	url: string;
}

export interface PatchSet {
	id: string;
	order: number;
	try_jobs: TryJob[] | null;
}

export interface ChangeListSummaryResponse {
	cl: ChangeList;
	patch_sets: PatchSet[] | null;
	num_total_patch_sets: number;
}

export interface TriageHistory {
	user: string;
	ts: string;
}

export interface Trace {
	label: TraceID;
	data: number[] | null;
	params: { [key: string]: string };
	comment_indices: number[] | null;
}

export interface DigestStatus {
	digest: Digest;
	status: Label;
}

export interface TraceGroup {
	tileSize: number;
	traces: Trace[] | null;
	digests: DigestStatus[] | null;
	total_digests: number;
}

export interface SRDiffDigest {
	numDiffPixels?: number;
	pixelDiffPercent?: number;
	maxRGBADiffs?: number[];
	dimDiffer?: boolean;
	diffs?: { [key: string]: number };
	digest: Digest;
	status: Label;
	paramset: ParamSet;
	n: number;
}

export interface SearchResult {
	digest: Digest;
	test: TestName;
	status: Label;
	triage_history: TriageHistory[] | null;
	paramset: ParamSet;
	traces: TraceGroup;
	refDiffs: { [key: string]: SRDiffDigest | null };
	closestRef: RefClosest;
}

export interface Commit {
	commit_time: number;
	hash: string;
	author: string;
	message: string;
	cl_url: string;
}

export interface TraceComment {
	id: ID;
	created_by: string;
	updated_by: string;
	created_ts: string;
	updated_ts: string;
	text: string;
	query: ParamSet;
}

export interface SearchResponse {
	digests: (SearchResult | null)[] | null;
	offset: number;
	size: number;
	commits: Commit[] | null;
	trace_comments: TraceComment[] | null;
	bulk_triage_data: TriageRequestData;
}

export interface TriageRequest {
	testDigestStatus: TriageRequestData;
	changelist_id: string;
	crs: string;
	imageMatchingAlgorithm: string;
}

export interface GUICorpusStatus {
	name: string;
	ok: boolean;
	minCommitHash: string;
	untriagedCount: number;
	negativeCount: number;
}

export interface StatusResponse {
	ok: boolean;
	firstCommit: Commit;
	lastCommit: Commit;
	totalCommits: number;
	filledCommits: number;
	corpStatus: (GUICorpusStatus | null)[] | null;
}

export type ParamSetResponse = { [key: string]: string[] | null };

export type Digest = string;

export type TestName = string;

export type Label = "untriaged" | "positive" | "negative";

export type ParamSet = { [key: string]: string[] | null };

export type TraceID = string;

export type RefClosest = "pos" | "neg" | "";

export type ID = string;

export type TriageRequestData = { [key: string]: { [key: string]: Label } };
