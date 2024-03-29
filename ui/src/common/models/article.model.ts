import { DateTime } from "luxon";

export interface ArticleGetParams {
  pbs?: string[];
  frm?: number;
  to?: number;
  ord: Ordering;
  srt: ArticleSortBy;
  sbj?: Subjects;
  oft: number;
  lmt: 10 | 25 | 50;
  pnd?: boolean;
}

export interface ArticleUpdateBody {
  description: string;
  subject: Subjects;
  accepted: boolean;
}

export interface Pagination {
  currentPage: number;
  pageSize: 10 | 25 | 50;
}

export interface ArticleFilters {
  sortBy: ArticleSortBy;
  order: Ordering;
  subject: Subjects;
  publishers: string[];
  dRange: [Date, Date] | [];
  pending?: boolean;
}

export interface Article {
  id: string;
  publisher: {
    id: string;
    name: string;
  };
  datePublished: string;
  dateRetrieved: string;
  title: string;
  description: string;
  link: string;
  type: Subjects;
  state: ArticleReviewState;
  matchedTerms: string[];
}

export interface ArticleDateSpan {
  max: string;
  min: string;
}

export enum Subjects {
  ALL = "",
  GHOST = "ghost",
  UFO = "ufo",
  CRYPTID = "cryptid",
}

export enum ArticleReviewState {
  ACCEPTED = "accepted",
  REVIEW = "review",
  REJECTED = "rejected",
}

export interface SubjectSelection {
  title: string;
  value: Subjects | "";
}

export enum ArticleSortBy {
  DATE = "date",
  TITLE = "title",
  DESCRIPTION = "description",
}

export enum Ordering {
  ASCENDING = "asc",
  DESCENDING = "desc",
}

export const defaultDateSpan: () => ArticleDateSpan = (): ArticleDateSpan => ({
  max: DateTime.now().plus({ month: 1 }).toISO() as string,
  min: DateTime.fromObject({ year: 1990 }).toISO() as string,
});
